
import { ref } from 'vue'
import type { EmailDocument } from '@/types/emails'

const zincHost = `http://localhost:4080/api`
const username = `admin` // Asegúrate de que esto esté definido en tu archivo de configuración
const password = `Complexpass#123`

export const useZincService = () => {
  const emails = ref<EmailDocument[]>([]) // Cambia 'any' por el tipo adecuado si lo conoces
  const error = ref<string | null>(null)

  const fetchEmails = async (userEmail: string) => {
    try {
      const response = await fetch(`${zincHost}/enron_emails/_search`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Basic ' + btoa(`${username}:${password}`),
        },
        body: JSON.stringify({
          search_type: "match",
          query: {
            term: userEmail,
            field: "to"
          },
          from: 0,
          max_results: 10
        }),
      })

      if (!response.ok) {
        throw new Error(`Error: ${response.status} ${response.statusText}`)
      }

      const data = await response.json()
      emails.value = data.hits.hits.map((hit: any) => hit._source) // Cambia 'any' por el tipo adecuado si lo conoces
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Error fetching emails'
      console.error("Error fetching emails:", error.value)
    }
  }

  return { emails, error, fetchEmails }
}