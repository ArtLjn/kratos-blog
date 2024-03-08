import axios from "axios";
var ipAdd;
export function getIp() {
    return new Promise((resolve, reject) => {
      let message;
      axios
        .get("https://api.ipify.org?format=json")
        .then((response) => {
          const ip = response.data.ip;
          ipAdd = ip;
          axios
            .get(`https://ipapi.co/${ip}/json/`)
            .then((response) => {
              message = response.data;
              resolve(message); // 使用resolve传递message
            });
        })
        .catch((error) => {
          reject(error);
        });
    });
  }
