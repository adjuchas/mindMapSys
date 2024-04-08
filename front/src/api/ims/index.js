import {requestIMS} from "@/api/request.js";
import {APPID, AUTHURI, SOURCE, UIDNAME} from "@/api/global.js";
import {sha256} from "js-sha256";


export function sendRedirect(uid) {
  let encryptedUid = sha256(uid).toString()
  let url = AUTHURI + "/authorize"
  let redirectUri =
    url +
    "?appId=" +
    APPID +
    "&encryptedUid=" +
    encryptedUid +
    "&source=" +
    SOURCE;
  window.location.href = redirectUri;
}

export function revoke(uid, accessToken){
  requestIMS({
    headers: {
      'Content-Type': "application/x-www-form-urlencoded",
    },
    url: "/revokeToken",
    method: "POST",
    body: "appId=" + APPID + "&uid=" + uid + "&accessToken=" + accessToken
  }).then(res => {
    let response = JSON.parse(res.data)
    console.log(response.toString())
  })
}



