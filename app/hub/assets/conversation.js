const urlParams = new URLSearchParams(window.location.search);
let eaddr = ""
const eaddrparam = urlParams.get("owner");
if (eaddrparam) {
  eaddr = eaddrparam;
}

let bucket = ""
const bucketparam = urlParams.get("bucket");
if (bucketparam) {
  bucket = bucketparam;
}

const itemsPerPage = 12;
let currentPage = 1;

function renderPage() {
  const cardContainer = document.getElementById("cardContainer");
  // Clear the current content
  cardContainer.innerHTML = '';

  const startIndex = (currentPage - 1) * itemsPerPage;
  listConversation(eaddr, startIndex, itemsPerPage)
    .then((data) => {
      data.forEach((conversationMeta) => {
        const card = createCard(conversationMeta);
        cardContainer.appendChild(card);
      });
    })
    .catch((error) => console.error("Error fetching conversation:", error));

  // Update the page number
  document.getElementById('pageNumber').innerText = `Page: ${currentPage}`;
}

// Handle "Previous" button click
document.getElementById('prevBtn').addEventListener('click', function () {
  if (currentPage > 1) {
    currentPage--;
    renderPage();
  }
});

// Handle "Next" button click
document.getElementById('nextBtn').addEventListener('click', function () {
  currentPage++;
  renderPage();
});

// Initial page render
renderPage();

function createCard(meta) {
  const card = document.createElement("div");
  card.className = "card";

  const h2 = document.createElement("h2");
  h2.textContent = meta.Name;
  card.appendChild(h2);

  const cardContent = document.createElement("div");
  cardContent.className = "card-content";
  card.appendChild(cardContent);

  const p5 = document.createElement("p");
  p5.setAttribute('data-label', 'Owner:');
  p5.textContent = meta.Owner;
  cardContent.appendChild(p5);

  const p6 = document.createElement("p");
  p6.setAttribute('data-label', 'Agent:');
  p6.textContent = meta.Bucket;
  cardContent.appendChild(p6);

  const p21 = document.createElement("p");
  p21.setAttribute('data-label', 'Creation:');
  p21.textContent = meta.CreatedAt;
  cardContent.appendChild(p21);

  const p3 = document.createElement("p");
  p3.setAttribute('data-label', 'Memories:');
  const needleLink = document.createElement('a');
  needleLink.href = 'needle.html?owner=' + meta.Owner + '&conversation=' + meta.Name;
  needleLink.textContent = '[click to show]';
  p3.appendChild(needleLink);
  cardContent.appendChild(p3);

  return card;
}

function listConversation(eaddr, offset, length) {
  return fetch(`/api/listConversation?owner=${eaddr}&bucket=${bucket}&offset=${offset}&length=${length}`)
    .then((response) => response.json())
    .then((data) => {
      return data;
    })
    .catch((error) => console.error("Error fetching needles:", error))
}

function search() {
  const searchInput = document.getElementById("searchInput").value;
  fetch(`/api/getConversation?owner=${eaddr}&bucket=${bucket}&name=${searchInput}`)
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
      cardContainer.appendChild(card);
    });
  }
}


