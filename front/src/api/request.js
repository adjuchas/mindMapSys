import axios from "axios";
import {BASEURI, AUTHURI} from "@/api/global.js";

export function requestBackend(config) {
  const instance = axios.create({
    baseURL: BASEURI,
    timeout: 3000,
    headers: {
      'Content-Type': 'application/json',
    }
  })
  return instance(config)
}

export function requestIMS(config) {
  const instance = axios.create({
    baseURL: AUTHURI,
    timeout: 3000,
    headers: {
      'Content-Type': "application/x-www-form-urlencoded",
    }
  })
  return instance(config)
}