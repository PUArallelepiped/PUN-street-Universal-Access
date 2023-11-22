import { goto } from "$app/navigation"
import { backendPath } from "./env"


export async function validateToken() {
    const response = await fetch(backendPath + '/validate', {
        method: 'GET',
        credentials: 'include',
    })
    if (response.status !== 200) {
        goto('/login')
    }
}