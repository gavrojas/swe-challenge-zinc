
import { ref } from 'vue'
import type { EmailDocument } from '@/types'

export const useZincService = () => {
  const emails = ref<EmailDocument[]>([])
  const error = ref<string | null>(null)

  return { emails, error }
}