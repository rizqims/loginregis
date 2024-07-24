const loginUrl = "http://localhost:8080/api/v1/users/login";

const loginData = {
  username: "anis",
  password: "admin",
};

// uhhh coba simplify
// async function loginUser(data) {
//   try {
//     const response = await fetch(loginUrl, {
//       method: "POST",
//       headers: {
//         "Content-Type": "application/json",
//       },
//       body: JSON.stringify(data),
//     });

//     if (!response.ok) {
//       throw new Error("Network response was not ok " + response.statusText);
//     }

//     const responseData = await response.json();
//     console.log("Login Response:", responseData);
//   } catch (error) {
//     console.error("There has been a problem with your fetch operation:", error);
//   }
// }

// loginUser(loginData);


async function loginUser(data) {
  fetch(loginUrl,
    {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data)
    }
  )
    .then(response => response.json())
    .then(data => {
      console.log(JSON.stringify(data));
    })
    .catch(error => console.error('Error fetching data:', error));
}

loginUser(loginData);