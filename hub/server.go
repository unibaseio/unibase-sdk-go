package hub

import (
	"net/http"
	"sync"

	"github.com/MOSSV2/dimo-sdk-go/lib/log"
	"github.com/MOSSV2/dimo-sdk-go/lib/logfs"
	"github.com/MOSSV2/dimo-sdk-go/lib/piece"
	"github.com/MOSSV2/dimo-sdk-go/lib/repo"
	"github.com/MOSSV2/dimo-sdk-go/lib/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

var logger = log.Logger("hub")

type Server struct {
	Router *gin.Engine

	typ string

	rp repo.Repo

	ps types.IPieceStore

	sync.RWMutex
	fscnt uint32
	lfs   map[string]*logfs.LogFS

	local common.Address

	auth types.Auth
}

func NewServer(rp repo.Repo) (*http.Server, error) {
	log.SetLogLevel("DEBUG")

	gin.SetMode(gin.ReleaseMode)

	localAddr := rp.Key().Address()

	logger.Infof("hub %s starting...", localAddr)

	router := gin.Default()

	auth, err := rp.Key().BuildAuth([]byte("hub"))
	if err != nil {
		return nil, err
	}

	s := &Server{
		Router: router,

		typ:   types.DownloaderType,
		local: localAddr,
		rp:    rp,
		ps:    piece.New(rp.MetaStore(), rp.DataStore()),
		auth:  auth,

		lfs: make(map[string]*logfs.LogFS),
	}

	s.load()
	go s.uploadTo()

	s.registRoute()

	srv := &http.Server{
		Addr:    rp.Config().API.Endpoint,
		Handler: s.Router,
	}

	return srv, nil
}

func (s Server) registRoute() {
	r := s.Router.Group("/api")

	s.addInfo(r)
	s.addDownload(r)
	s.addUpload(r)
	s.addUploadData(r)
}

func (s Server) addInfo(g *gin.RouterGroup) {
	g.Group("/").GET("/info", func(c *gin.Context) {
		res := types.EdgeReceipt{
			EdgeMeta: types.EdgeMeta{
				Type: s.typ,
				Name: s.local,
			},
		}

		c.JSON(http.StatusOK, res)
	})
}