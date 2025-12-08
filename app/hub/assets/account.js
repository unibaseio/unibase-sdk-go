const itemsPerPage = 12;
let currentPage = 1;

document.addEventListener("DOMContentLoaded", () => {
  // 初始化页面
  renderPage();

  // 添加分页按钮事件监听
  document.getElementById('prevBtn').addEventListener('click', function () {
    if (currentPage > 1) {
      currentPage--;
      renderPage();
    }
  });

  document.getElementById('nextBtn').addEventListener('click', function () {
    currentPage++;
    renderPage();
  });
});

function renderPage() {
  const cardContainer = document.getElementById("cardContainer");
  // 清空当前内容
  cardContainer.innerHTML = '';

  const startIndex = (currentPage - 1) * itemsPerPage;
  listAccount(startIndex, itemsPerPage)
    .then((data) => {
      if (data.length === 0 && currentPage > 1) {
        // 如果当前页没有数据且不是第一页，返回上一页
        currentPage--;
        renderPage();
        return;
      }
      data.forEach((meta) => {
        const card = createCard(meta);
        cardContainer.appendChild(card);
      });
    })
    .catch((error) => console.error("Error fetching accounts:", error));

  // 更新页码显示
  document.getElementById('pageNumber').innerText = `Page: ${currentPage}`;
}

function createCard(meta) {
  const card = document.createElement("div");
  card.className = "card";

  const h2 = document.createElement("h2");
  h2.textContent = meta.Name;
  card.appendChild(h2);

  const cardContent = document.createElement("div");
  cardContent.className = "card-content";
  card.appendChild(cardContent);

  const p0 = document.createElement("p");
  p0.setAttribute('data-label', 'Creation:');
  p0.textContent = meta.CreatedAt;
  cardContent.appendChild(p0);

  const p3 = document.createElement("p");
  p3.setAttribute('data-label', 'Agents:');
  const bucketLink = document.createElement('a');
  bucketLink.href = 'bucket.html?owner=' + meta.Name;
  bucketLink.textContent = '[click to show]';
  p3.appendChild(bucketLink);
  cardContent.appendChild(p3);

  const p4 = document.createElement("p");
  p4.setAttribute('data-label', 'Conversations:');
  const conversationLink = document.createElement('a');
  conversationLink.href = 'conversation.html?owner=' + meta.Name;
  conversationLink.textContent = '[click to show]';
  p4.appendChild(conversationLink);
  cardContent.appendChild(p4);

  const p1 = document.createElement("p");
  p1.setAttribute('data-label', 'Memories:');
  const needleLink = document.createElement('a');
  needleLink.href = 'needle.html?owner=' + meta.Name;
  needleLink.textContent = '[click to show]';
  p1.appendChild(needleLink);
  cardContent.appendChild(p1);

  const p2 = document.createElement("p");
  p2.setAttribute('data-label', 'Volumes :');
  const volumeLink = document.createElement('a');
  volumeLink.href = 'volume.html?owner=' + meta.Name;
  volumeLink.textContent = '[click to show]';
  p2.appendChild(volumeLink);
  cardContent.appendChild(p2);


  return card;
}

function listAccount(offset = 0, length = itemsPerPage) {
  return fetch(`/api/listAccount?offset=${offset}&length=${length}`)
    .then((response) => response.json())
    .then((data) => {
      return data;
    })
    .catch((error) => console.error("Error fetching accounts:", error))
}

function search() {
  const searchInput = document.getElementById("searchInput").value;
  fetch(`/api/getAccount?owner=${searchInput}`)
    .then((response) => response.json())
    .then((data) => {
      displayResults(data);
    })
    .catch((error) => console.error("Error:", error));
}

function displayResults(data) {
  const resultsElement = document.getElementById("cardContainer");
  resultsElement.innerHTML = ""; // 清空先前的结果

  if (data.length === 0) {
    resultsElement.innerText = "No results found.";
  } else {
    data.forEach((meta) => {
      const card = createCard(meta);
      resultsElement.appendChild(card);
    });
  }

  // 重置分页状态
  currentPage = 1;
  document.getElementById('pageNumber').innerText = `Page: ${currentPage}`;
}