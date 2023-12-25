import { redirect } from "@sveltejs/kit";
import { jwtDecode } from "jwt-decode";

export async function getId() {
    // if (typeof window === "undefined") {
    //     return 0;
    // }
    let cookie = document.cookie;
    if (cookie == "" || cookie == undefined) {
        throw new Error("No cookie");
    }
    interface token {
        id: string;
    }
    let decoded: token = jwtDecode(cookie.split("=")[1]);
    return decoded.id;
}

export async function getIdByToken(token: string) {
    interface token {
        id: string;
    }
    if (token == "" || token == undefined) {
        throw new Error("No cookie");
    }
    let decoded: token = jwtDecode(token);
    return decoded.id;
}