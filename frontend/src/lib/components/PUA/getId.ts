import { jwtDecode } from "jwt-decode";

export async function getId() {
    let cookie = document.cookie;
    interface token {
        id: string;
    }
    let decoded: token = jwtDecode(cookie.split("=")[1]);
    return decoded.id;
}
