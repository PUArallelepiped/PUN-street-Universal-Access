import { jwtDecode } from 'jwt-decode';

export async function getId() {
	// if (typeof window === "undefined") {
	//     return 0;
	// }
	const cookie = document.cookie;
	if (cookie == '' || cookie == undefined) {
		throw new Error('No cookie');
	}
	interface token {
		id: string;
	}
	const decoded: token = jwtDecode(cookie.split('=')[1]);
	return decoded.id;
}

export async function getIdByToken(token: string) {
	interface token {
		id: string;
	}
	if (token == '' || token == undefined) {
		throw new Error('No cookie');
	}
	const decoded: token = jwtDecode(token);
	return decoded.id;
}
