import { TicketcheckInterface } from "../../interface/ITicketcheck.ts";
// import { TicketInterface } from "../../interface/ITicket.ts";

const apiUrl = "http://localhost:8000";


//
async function Checkin(TicketID: TicketcheckInterface) {
  const requestOptions = {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      ticket_id: TicketID,
      time_stamp: new Date().toISOString(), // Example of generating a timestamp
      status: "Check in", // Replace with actual status
      
    }),
  };

  let res = await fetch(`${apiUrl}/ticketcheck/`, requestOptions)
    .then((response) => response.json())
    .then((res) => {
      if (res.message =="check in") {
        return { status: true, message: "complete" };
      } else {
        return { status: false, message: "failed" };
      }
    })
    .catch((error) => {
      return { status: false, message: "failed" };
    });

  return res;
}

//

async function GetTicketcheck() {
  const requestOptions = {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  };


  let res = await fetch(`${apiUrl}/ticketcheck`, requestOptions)
    .then((response) => response.json())
    .then((res) => {
      if (res.data) {
        return res.data;
      } else {
        return false;
      }
    });

  return res;
}


async function DeleteTicketcheckByID(id: Number | undefined) {
  const requestOptions = {
    method: "DELETE",
    headers: {
      "Content-Type": "application/json",
    },
  };

  let res = await fetch(`${apiUrl}/ticketcheck/${id}`, requestOptions)
    .then((response) => response.json())
    .then((res) => {
      if (res.data) {
        return res.data;
      } else {
        return false;
      }
    });

  return res;
}

async function GetTicketchecktById(id: Number | undefined) {
  const requestOptions = {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  };

  let res = await fetch(`${apiUrl}/ticketcheck/${id}`, requestOptions)
    .then((response) => response.json())
    .then((res) => {
      if (res.data) {
        return res.data;
      } else {
        return false;
      }
    });

  return res;
}

//สร้างข้อมูล ticketcheck


// async function UpdateStudent(data: TicketcheckInterface) {
//   const requestOptions = {
//     method: "PATCH",
//     headers: { 
//       "Content-Type": "application/json" 
//     },
//     body: JSON.stringify(data),
//   };

//   let res = await fetch(`${apiUrl}/students`, requestOptions)
//     .then((response) => response.json())
//     .then((res) => {
//       if (res.data) {
//         return { status: true, message: res.data };
//       } else {
//         return { status: false, message: res.error };
//       }
//     });

//   return res;
// }

export {
  Checkin,

  GetTicketcheck,
  DeleteTicketcheckByID,
  GetTicketchecktById,
//   UpdateStudent,
};