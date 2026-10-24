const BASE_URL = ""

const get_many_objects = document.querySelector("#get_many_objects")
const get_one_object = document.querySelector("#get_one_object")
const post_one_object = document.querySelector("#post_one_object")
const put_one_object = document.querySelector("#put_one_object")
const patch_one_object = document.querySelector("#patch_one_object")
const delete_one_object  = document.querySelector("#delete_one_object")

let obj_id

// get_many_objects.addEventListener('click',() => {
//   fetch(`${BASE_URL}/objects`).then(response => {
//     return response.json();
//   }).then(data => {
//     console.log(data.map(obj => obj.name))
//   })
// });

// get_one_object.addEventListener('click',() => {
//   fetch(`${BASE_URL}/objects/${obj_id}`).then(response => {
//     return response.json();
//   }).then(data => {
//     if (data === undefined){
//       console.error("id with that data not found")
//     }else{
//     console.log(data.name)
//     }
//   })
// });

// post_one_object.addEventListener('click',() => {
//   fetch(`${BASE_URL}/objects`,{
//     method: 'POST',
//     headers: {"Content-Type": "application/json"},
//     body: JSON.stringify({
//     "name": "Apple MacBook Pro 16123",
//     "data": {
//       "year": 2019,
//       "price": 1849.99,
//       "CPU model": "Intel Core i9",
//       "Hard disk size": "1 TB"
//       }
//     })
//   }).then(response => {
//     return response.json();
//   }).then(data => {
//     console.log(data)
//     obj_id = data.id
//   })
// });

// put_one_object.addEventListener('click',() => {
//   fetch(`${BASE_URL}/objects/${obj_id}`,{
//     method: 'PUT',
//     headers: {"Content-Type": "application/json"},
//     body: JSON.stringify({
//     "name": "Apple MacBook Pro 16123789",
//     "data": {
//       "year": 2019,
//       "price": 1849.99,
//       "CPU model": "Intel Core i9",
//       "Hard disk size": "1 TB"
//       }
//     })
//   }).then(response => {
//     return response.json();
//   }).then(data => {
//     console.log(data)
//   })
// });

// patch_one_object.addEventListener('click',() => {
//   fetch(`${BASE_URL}/objects/${obj_id}`,{
//     method: 'PATCH',
//     headers: {"Content-Type": "application/json"},
//     body: JSON.stringify({
//   "name": "Apple MacBook Pro 16 (Updated Name)"
// })
//   }).then(response => {
//     return response.json();
//   }).then(data => {
//     console.log(data)
//   })
// });

// delete_one_object.addEventListener('click',() => {
//   fetch(`${BASE_URL}/objects/${obj_id}`,{
//     method: 'DELETE',
//     headers: {"Content-Type": "application/json"}
//   }).then(response => {
//     return response.json();
//   }).then(data => {
//     console.log(data)
//     obj_id = 1
//   })
// });

// ------------------------------------

get_many_objects.addEventListener('click', async () => {
  try {
    const response = await fetch(`${BASE_URL}/objects`);
    const data = await response.json();
    console.log(data.map(obj => obj.name));
  } catch (e) {
    console.error(e);
  }
});


get_one_object.addEventListener('click',async () => {
  try {
    const response = await fetch(`${BASE_URL}/objects/${obj_id}`);
    const data = await response.json();
    if (data === undefined){
      console.error("id with that data not found")
    }else{
    console.log(data.name)
    }
  } catch (e) {
    console.error(e);
  }
});

post_one_object.addEventListener('click',async () => {
  try {
    const response = await fetch(`${BASE_URL}/objects`,{
      method: 'POST',
      headers: {"Content-Type": "application/json"},
      body: JSON.stringify({
      "name": "Apple MacBook Pro 16123",
      "data": {
        "year": 2019,
        "price": 1849.99,
        "CPU model": "Intel Core i9",
        "Hard disk size": "1 TB"
        }
      })
    });
    const data = await response.json();
    console.log(data)
    obj_id = data.id
  } catch (e) {
    console.error(e);
  }
});

put_one_object.addEventListener('click',async () => {
  try {
    const response = await fetch(`${BASE_URL}/objects/${obj_id}`,{
      method: 'PUT',
      headers: {"Content-Type": "application/json"},
      body: JSON.stringify({
      "name": "Apple MacBook Pro 16123789",
      "data": {
        "year": 2019,
        "price": 1849.99,
        "CPU model": "Intel Core i9",
        "Hard disk size": "1 TB"
        }
      })
    });
    const data = await response.json();
    console.log(data)
  } catch (e) {
    console.error(e);
  }
});

patch_one_object.addEventListener('click',async () => {
  try {
    const response = await fetch(`${BASE_URL}/objects/${obj_id}`,{
      method: 'PATCH',
      headers: {"Content-Type": "application/json"},
      body: JSON.stringify({
    "name": "Apple MacBook Pro 16 (Updated Name)"
      })
    });
    const data = await response.json();
    console.log(data)
  } catch (e) {
    console.error(e);
  }
});

delete_one_object.addEventListener('click', async () => {
  try {
    const response = await fetch(`${BASE_URL}/objects/${obj_id}`, {
      method: 'DELETE',
      headers: { "Content-Type": "application/json" }
    });
    
    if (response.ok) {
      // Safely check for empty response body instead of calling .json()
      const text = await response.text();
      const data = text ? JSON.parse(text) : { success: true };
      
      console.log(data);
      // obj_id = 1;
    } else {
      console.error(`Server error: ${response.status}`);
    }
  } catch (e) {
    console.error('Network or parsing error:', e);
  }
});