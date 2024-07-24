const loginUrl = "http://localhost:8080/api/v1/users/login";

// const data = "./data.json";

async function fetchData(url) {
  try {
    const response = await fetch(url);
    if (!response.ok) {
      throw new Error("Network response was not ok");
    }
    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Error fetching data:", error);
  }
}

async function regisUser(data) {
  fetch(loginUrl, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  })
    .then((response) => response.json())
    .then((data) => {
      console.log(JSON.stringify(data));
    })
    .catch((error) => console.error("Error fetching data:", error));
}

async function main() {
  const data = await fetchData("");
  await regisUser(data);
}

// Pertama, muat data dari file data.json
// fetch("./data.json")
//   .then((response) => {
//     if (!response.ok) {
//       throw new Error("Network response was not ok");
//     }
//     return response.json();
//   })
//   .then((data) => {
//     // Kirimkan data tersebut ke endpoint API
//     return fetch(loginUrl, {
//       method: "POST",
//       headers: {
//         "Content-Type": "application/json",
//       },
//       body: JSON.stringify(data),
//     });
//   })
//   .then((response) => {
//     if (!response.ok) {
//       console.log(response);
//       throw new Error("Network response was not ok");
//     }
//     response.json();
//   })
//   .then((data) => console.log("Success:", data))
//   .catch((error) => console.error("Error:", error));
