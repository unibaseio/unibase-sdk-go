package hub

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/MOSSV2/dimo-sdk-go/contract"
	lerror "github.com/MOSSV2/dimo-sdk-go/lib/error"
	"github.com/MOSSV2/dimo-sdk-go/lib/key"
	"github.com/MOSSV2/dimo-sdk-go/lib/logfs"
	"github.com/MOSSV2/dimo-sdk-go/lib/types"
	"github.com/MOSSV2/dimo-sdk-go/sdk"
	"github.com/gin-gonic/gin"
)

func (s *Server) addUploadData(g *gin.RouterGroup) {
	g.Group("/").POST("/uploadData", func(c *gin.Context) {
		addr := c.PostForm("owner")

		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(599, lerror.ToAPIError("hub", err))
			return
		}

		if file == nil {
			c.JSON(599, lerror.ToAPIError("hub", fmt.Errorf("file is nil")))
			return
		}

		fr, err := file.Open()
		if err != nil {
			c.JSON(599, lerror.ToAPIError("hub", err))
			return
		}

		if file.Size == 0 {
			c.JSON(599, lerror.ToAPIError("hub", fmt.Errorf("empty file")))
			return
		}
		mm, err := s.logFSWrite(addr, file.Filename, fr)
		if err != nil {
			c.JSON(599, lerror.ToAPIError("hub", err))
			return
		}

		c.JSON(http.StatusOK, mm)
	})
}

func (s *Server) addUpload(g *gin.RouterGroup) {
	g.Group("/").POST("/upload", func(c *gin.Context) {
		var mjson types.MemeStruct

		err := c.ShouldBindJSON(&mjson)
		if err != nil {
			c.JSON(599, lerror.ToAPIError("hub", err))
			return
		}

		var buf bytes.Buffer
		buf.WriteString(mjson.Message)

		mm, err := s.logFSWrite(mjson.Owner, mjson.ID, &buf)
		if err != nil {
			c.JSON(599, lerror.ToAPIError("hub", err))
			return
		}

		c.JSON(http.StatusOK, mm)
	})
}

func (s *Server) logFSWrite(addr string, key string, r io.Reader) (types.MemeMeta, error) {
	var err error
	if addr == "" {
		addr = s.local.String()
	}
	s.Lock()
	fs, ok := s.lfs[addr]
	if !ok {
		fspath := filepath.Join(s.rp.Path(), LOGFS)
		fs, err = logfs.New(s.rp.MetaStore(), fspath, addr)
		if err != nil {
			s.Unlock()
			return types.MemeMeta{}, err
		}
		s.lfs[addr] = fs
		logger.Infof("start new log inst: %s %d", addr, s.fscnt)

		dsKey := types.NewKey(types.DsLogFS, LOGINST, addr)
		has, err := s.rp.MetaStore().Has(dsKey)
		if err != nil || !has {
			buf := make([]byte, 4)
			binary.BigEndian.PutUint32(buf, 0)
			s.rp.MetaStore().Put(dsKey, buf)

			dsKey = types.NewKey(types.DsLogFS, LOGINST, s.fscnt)
			s.rp.MetaStore().Put(dsKey, []byte(addr))

			dsKey = types.NewKey(types.DsLogFS, LOGINST)
			s.fscnt++
			binary.BigEndian.PutUint32(buf, s.fscnt)
			s.rp.MetaStore().Put(dsKey, buf)
		}
	}
	s.Unlock()
	rbytes, err := io.ReadAll(r)
	if err != nil {
		return types.MemeMeta{}, err
	}

	err = fs.Put([]byte(key), rbytes)
	if err != nil {
		return types.MemeMeta{}, err
	}

	lm, err := fs.GetMeta([]byte(key))
	if err != nil {
		return types.MemeMeta{}, err
	}

	mm := types.MemeMeta{
		File:  fmt.Sprintf("%s-%d.log", addr, lm.Index),
		Start: lm.Start,
		Size:  lm.Size,
	}

	return mm, nil
}

func (s *Server) logFSRead(addr string, key string, w io.Writer) (int64, error) {
	var err error
	if addr == "" {
		addr = s.local.String()
	}

	s.Lock()
	fs, ok := s.lfs[addr]
	if !ok {
		fspath := filepath.Join(s.rp.Path(), LOGFS)
		fs, err = logfs.New(s.rp.MetaStore(), fspath, addr)
		if err != nil {
			s.Unlock()
			return 0, err
		}
		s.lfs[addr] = fs

		logger.Infof("start new log inst: %s %d", addr, s.fscnt)

		dsKey := types.NewKey(types.DsLogFS, LOGINST, addr)
		has, err := s.rp.MetaStore().Has(dsKey)
		if err != nil || !has {
			buf := make([]byte, 4)
			binary.BigEndian.PutUint32(buf, 0)
			s.rp.MetaStore().Put(dsKey, buf)

			dsKey = types.NewKey(types.DsLogFS, LOGINST, s.fscnt)
			s.rp.MetaStore().Put(dsKey, []byte(addr))

			dsKey = types.NewKey(types.DsLogFS, LOGINST)
			s.fscnt++
			binary.BigEndian.PutUint32(buf, s.fscnt)
			s.rp.MetaStore().Put(dsKey, buf)
		}
	}
	s.Unlock()

	lm, err := fs.GetMeta([]byte(key))
	if err != nil {
		return 0, err
	}

	wbytes, err := fs.GetData(lm)
	if err != nil {
		// todo: get from remote
		return 0, err
	}
	n, err := w.Write(wbytes)
	if err != nil {
		return 0, err
	}

	return int64(n), nil
}

func (s *Server) load() error {
	fspath := filepath.Join(s.rp.Path(), LOGFS)
	fs, err := logfs.New(s.rp.MetaStore(), fspath, s.local.String())
	if err != nil {
		return err
	}
	s.lfs[s.local.String()] = fs

	dsKey := types.NewKey(types.DsLogFS, LOGINST)
	val, err := s.rp.MetaStore().Get(dsKey)
	if err == nil && len(val) == 4 {
		s.fscnt = binary.BigEndian.Uint32(val)
	}

	if s.fscnt == 0 {
		s.fscnt = 1
		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, s.fscnt)
		s.rp.MetaStore().Put(dsKey, buf)

		dsKey := types.NewKey(types.DsLogFS, LOGINST, 0)
		s.rp.MetaStore().Put(dsKey, []byte(s.local.String()))
	}

	logger.Infof("load log inst: %d", s.fscnt)
	return nil
}

func (s *Server) uploadTo() {
	sk := s.rp.Key().Export().PrivateKey
	au, err := key.BuildAuth(sk, []byte("upload"))
	if err != nil {
		panic(err)
	}

	sdk.Login(sdk.ServerURL, au)

	policy := types.Policy{
		N: 6,
		K: 4,
	}

	for {
		time.Sleep(time.Minute)
		err := contract.CheckBalance(au.Addr)
		if err != nil {
			time.Sleep(time.Minute)
			continue
		}

		for i := uint32(0); i < s.fscnt; i++ {
			dsKey := types.NewKey(types.DsLogFS, LOGINST, i)
			val, err := s.rp.MetaStore().Get(dsKey)
			if err != nil {
				break
			}

			key := string(val)

			logger.Debugf("check: %s %d", key, i)
			dsKey = types.NewKey(types.DsLogFS, key)
			val, err = s.rp.MetaStore().Get(dsKey)
			if err != nil || len(val) != 4 {
				continue
			}
			curIndex := binary.BigEndian.Uint32(val)

			next := uint32(0)
			dsKey = types.NewKey(types.DsLogFS, LOGINST, key)
			val, err = s.rp.MetaStore().Get(dsKey)
			if err == nil && len(val) == 4 {
				next = binary.BigEndian.Uint32(val)
			}

			logger.Debugf("check: %s %d %d", key, next, curIndex)
			if next >= curIndex {
				continue
			}

			for i := next; i < curIndex; i++ {
				fname := fmt.Sprintf("%s-%d.log", key, i)
				fp := filepath.Join(s.rp.Path(), LOGFS, key, fmt.Sprintf("%d.log", i))

				fr, err := sdk.GetFileReceipt(sdk.ServerURL, au, fname)
				if err == nil {
					logger.Infof("%s-%d.log is already uploaded, check its piece onchain", key, i)
					er, err := sdk.ListEdge(sdk.ServerURL, au, types.StreamType)
					if err != nil {
						break
					}
					suc := 0
					for _, pn := range fr.Pieces {
						for _, st := range er.Edges {
							pr, err := sdk.GetPieceReceipt(st.ExposeURL, au, pn)
							if err == nil {
								if pr.Serial > 0 {
									suc++
								} else {
									err = contract.AddPiece(sk, pr.PieceCore)
									if err == nil {
										suc++
									}
								}
							}
						}
					}
					if suc == len(fr.Pieces) {
						buf := make([]byte, 4)
						binary.BigEndian.PutUint32(buf, i+1)
						s.rp.MetaStore().Put(dsKey, buf)
						continue
					}
					break
				}
				// upload to stream and submit to gateway
				res, streamer, err := sdk.Upload(sdk.ServerURL, au, policy, fp, fname)
				if err != nil {
					break
				}
				pcs, err := sdk.CheckFileFull(res, streamer, fp)
				if err != nil {
					break
				}
				log.Printf("upload %s to %s, sha256: %s\n", fp, streamer, res.Hash)
				log.Printf("submit %s to chain\n", res.Name)
				// submit meta to chain
				for _, pc := range pcs {
					err = contract.AddPiece(sk, pc)
					if err != nil {
						break
					}
				}
				if err != nil {
					break
				}
				buf := make([]byte, 4)
				binary.BigEndian.PutUint32(buf, i+1)
				s.rp.MetaStore().Put(dsKey, buf)
			}
		}
	}
}