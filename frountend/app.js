const API = "http://localhost:8080";

async function addExpense() {
  const key = crypto.randomUUID();

  const body = {
    amount: Number(document.getElementById("amount").value),
    category: document.getElementById("category").value,
    description: document.getElementById("description").value,
    date: document.getElementById("date").value
  };

  await fetch(API + "/expenses", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      "Idempotency-Key": key
    },
    body: JSON.stringify(body)
  });

  loadExpenses();
}

async function loadExpenses() {
  const category = document.getElementById("filter").value;
  const res = await fetch(API + "/expenses?category=" + category);
  const data = await res.json();

  const list = document.getElementById("list");
  list.innerHTML = "";

  let total = 0;

  data.forEach(e => {
    total += e.amount;
    const li = document.createElement("li");
    li.innerText = `${e.category} - ₹${e.amount / 100}`;
    list.appendChild(li);
  });

  document.getElementById("total").innerText =
    `Total: ₹${total / 100}`;
}

loadExpenses();
