const urlParams = new URLSearchParams(window.location.search);
let eaddr = ""
const eaddrparam = urlParams.get("owner");
if (eaddrparam) {
  eaddr = eaddrparam;
}

const itemsPerPage = 12;
let currentPage = 1;

function renderPage() {
  const cardContainer = document.getElementById("cardContainer");
  // Clear the current content
  cardContainer.innerHTML = '';

  const startIndex = (currentPage - 1) * itemsPerPage;
  listBucket(eaddr, startIndex, itemsPerPage)
    .then((data) => {
      data.forEach((fileMeta) => {
        const card = createCard(fileMeta);
        cardContainer.appendChild(card);
      });
    })
    .catch((error) => console.error("Error fetching bucket:", error));

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

  const p21 = document.createElement("p");
  p21.setAttribute('data-label', 'Creation:');
  p21.textContent = meta.CreatedAt;
  cardContent.appendChild(p21);

  if (meta.Description) {
    const p22 = document.createElement("p");
    p22.setAttribute('data-label', 'Description:');

    if (meta.Description.length > 40) {
      const shortText = meta.Description.substring(0, 40);
      const expandBtn = document.createElement("span");
      expandBtn.textContent = "...";
      expandBtn.style.cursor = "pointer";
      expandBtn.style.color = "#007bff";
      expandBtn.style.textDecoration = "underline";

      p22.textContent = shortText;
      p22.appendChild(expandBtn);

      let isExpanded = false;
      expandBtn.addEventListener('click', function () {
        if (isExpanded) {
          p22.textContent = shortText;
          p22.appendChild(expandBtn);
          isExpanded = false;
        } else {
          p22.textContent = meta.Description;
          isExpanded = true;
        }
      });
    } else {
      p22.textContent = meta.Description;
    }

    cardContent.appendChild(p22);
  }

  if (meta.Transport) {
    const p23 = document.createElement("p");
    p23.setAttribute('data-label', 'Transport:');
    p23.textContent = meta.Transport;
    cardContent.appendChild(p23);
  }

  if (meta.State) {
    const p24 = document.createElement("p");
    p24.setAttribute('data-label', 'State:');
    p24.textContent = meta.State;
    cardContent.appendChild(p24);
  }

  if (meta.Last) {
    const p25 = document.createElement("p");
    p25.setAttribute('data-label', 'Last:');
    p25.textContent = meta.Last;
    cardContent.appendChild(p25);
  }

  const p4 = document.createElement("p");
  p4.setAttribute('data-label', 'Conversations:');
  const conversationLink = document.createElement('a');
  conversationLink.href = 'conversation.html?owner=' + meta.Owner + '&bucket=' + meta.Name;
  conversationLink.textContent = '[click to show]';
  p4.appendChild(conversationLink);
  cardContent.appendChild(p4);

  const p3 = document.createElement("p");
  p3.setAttribute('data-label', 'Memories:');
  const needleLink = document.createElement('a');
  needleLink.href = 'needle.html?owner=' + meta.Owner + '&bucket=' + meta.Name;
  needleLink.textContent = '[click to show]';
  p3.appendChild(needleLink);
  cardContent.appendChild(p3);

  return card;
}

function listBucket(eaddr, offset, length) {
  return fetch(`/api/listBucket?owner=${eaddr}&offset=${offset}&length=${length}`)
    .then((response) => response.json())
    .then((data) => {
      return data;
    })
    .catch((error) => console.error("Error fetching needles:", error))
}

function search() {
  const searchInput = document.getElementById("searchInput").value;
  fetch(`/api/getBucket?bucket=${searchInput}`)
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


