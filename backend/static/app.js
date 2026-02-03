const API = "";

let allExpenses = [];
let allCategories = [];
let sortDesc = true;

/* ---------------- ADD EXPENSE ---------------- */

async function addExpense() {
  const status = document.getElementById("status");
  const btn = document.getElementById("submitBtn");

  const amount = Number(document.getElementById("amount").value);
  const category = document.getElementById("category").value.trim();
  const description = document.getElementById("description").value;
  const date = document.getElementById("date").value;

  // Validation
  if (amount <= 0 || !category || !date) {
    status.textContent = "Amount, category and date are required";
    status.className = "error";
    return;
  }

  // Idempotency for retries / refresh
  const key =
    localStorage.getItem("lastKey") || crypto.randomUUID();
  localStorage.setItem("lastKey", key);

  btn.disabled = true;
  status.textContent = "Saving...";
  status.className = "loading";

  try {
    await fetch(`${API}/expenses`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Idempotency-Key": key
      },
      body: JSON.stringify({
        amount: amount * 100, // ₹ → paise
        category,
        description,
        date
      })
    });

    localStorage.removeItem("lastKey");

    // Clear inputs after successful add
    document.getElementById("amount").value = "";
    document.getElementById("category").value = "";
    document.getElementById("description").value = "";
    document.getElementById("date").value = "";

    status.textContent = "Expense added";
    status.className = "";

    loadExpenses();
  } catch {
    status.textContent = "Failed to save expense";
    status.className = "error";
  } finally {
    btn.disabled = false;
  }
}

/* ---------------- LOAD EXPENSES ---------------- */

async function loadExpenses() {
  const table = document.getElementById("expenseTable");
  const totalEl = document.getElementById("total");

  table.innerHTML = "";
  totalEl.textContent = "Loading...";

  try {
    const res = await fetch(`${API}/expenses`);
    allExpenses = await res.json();

    // Extract unique categories
    allCategories = [...new Set(allExpenses.map(e => e.category))];

    renderTable(allExpenses);
  } catch {
    totalEl.textContent = "Failed to load expenses";
  }
}

/* ---------------- RENDER TABLE ---------------- */

function renderTable(expenses) {
  const table = document.getElementById("expenseTable");
  const totalEl = document.getElementById("total");

  table.innerHTML = "";

  let total = 0;

  expenses
expenses
  .sort((a, b) =>
    sortDesc
      ? b.date.localeCompare(a.date)
      : a.date.localeCompare(b.date)
  )
    .forEach(e => {
      total += e.amount;

      const row = document.createElement("tr");
      row.innerHTML = `
        <td>${(e.amount / 100).toFixed(2)}</td>
        <td>${e.category}</td>
        <td>${e.description}</td>
        <td>${e.date}</td>
      `;
      table.appendChild(row);
    });

  totalEl.textContent = `Total: ₹${(total / 100).toFixed(2)}`;
}

/* ---------------- CATEGORY FILTER ---------------- */

function onFilterInput() {
  const input = document
    .getElementById("filterInput")
    .value
    .toLowerCase();

  const dropdown = document.getElementById("categoryDropdown");
  dropdown.innerHTML = "";

  if (!input) {
    dropdown.style.display = "none";
    return;
  }

  const matches = allCategories
    .filter(c => c.toLowerCase().startsWith(input))
    .slice(0, 5);

  if (matches.length === 0) {
    dropdown.style.display = "none";
    return;
  }

  matches.forEach(cat => {
    const option = document.createElement("option");
    option.value = cat;
    option.textContent = cat;
    dropdown.appendChild(option);
  });

  dropdown.style.display = "block";
}


function onCategorySelect() {
  const dropdown = document.getElementById("categoryDropdown");
  const selected = dropdown.value;

  document.getElementById("filterInput").value = selected;
  dropdown.style.display = "none";

  const filtered = allExpenses.filter(
    e => e.category === selected
  );

  renderTable(filtered);
}


function clearFilter() {
  document.getElementById("filterInput").value = "";
  document.getElementById("categoryDropdown").style.display = "none";
  document.getElementById("categoryDropdown").innerHTML = "";
  renderTable(allExpenses);
}


/* ---------------- INIT ---------------- */

loadExpenses();
function toggleSort() {
  sortDesc = !sortDesc;

  // Apply sort on currently visible data
  const input = document.getElementById("filterInput").value;

  if (input) {
    const filtered = allExpenses.filter(
      e => e.category === input
    );
    renderTable(filtered);
  } else {
    renderTable(allExpenses);
  }
}
