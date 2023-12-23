import { jwtDecode } from "jwt-decode";

export async function getId() {
    if (typeof window === "undefined") {
        return 0;
    }
    let cookie = document.cookie;
    interface token {
        id: string;
    }
    let decoded: token = jwtDecode(cookie.split("=")[1]);
    return decoded.id;
}
