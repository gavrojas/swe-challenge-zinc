import type { OptionsApiCall } from '@/types/index'
import { useAuthStore } from '@/stores/auth'
import type { EmailDocument } from '@/types'
import { ref } from 'vue'

const BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080/emails'

type ApiCallOptions = OptionsApiCall | undefined
export const useZincService = () => {
  const emails = ref<EmailDocument[]>([])
  const error = ref<string | null>(null)

  async function apiHandleEmails(
    path: string,
    { method = 'GET', data, headers, notifyLogout }: ApiCallOptions = {},
  ) {
    
    const sessionToken = sessionStorage.getItem('token')
    const requestHeaders = {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${sessionToken}`,
      ...headers,
    }
    const body = method !== 'GET' && data ? JSON.stringify(data) : undefined
    try {
      const response = await fetch(BASE_URL + path, {
        method,
        headers: requestHeaders,
        body,
      })

      if(notifyLogout){
        await fetch(BASE_URL + '/auth/logout',{
          method: 'POST',
          headers: {
            'Authorization': `Bearer ${sessionToken}`,
            'Content-Type': 'application/json'
          },
        })}

      if (!response.ok) {
        // Si el error es de sesión expirada (401), redirigir al login
        if (response.status === 401) {
          const authStore = useAuthStore()
          authStore.clearSession() 
          throw new Error('Session expired. Redirecting to login...')
        }

        throw new Error(`Error: ${response.status} ${response.statusText}`)
      }

      if (method === 'DELETE') {
        return response.ok
      }

      const data = await response.json()
      return data.hits.hits.map((hit: any) => hit._source)
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Error fetching emails'
      console.error("Error fetching emails:", error.value)
    }
  }

  return { emails, error, apiHandleEmails }
}