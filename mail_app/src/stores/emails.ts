import { defineStore } from 'pinia';
import type { EmailDocument } from '@/types';
import { useZincService } from '@/services/utils'

const { apiHandleEmails } = useZincService()

export const useEmailStore = defineStore('email', {
    state: () => ({
        emails: [] as EmailDocument[],
        error: null as string | null,
        folders: [
            'all_documents',
            'inbox',
            'sent',
            'sent_items',
            'deleted_items',
            'discussion_threads',
            'tasks',
            'calendar',
            'contacts',
        ],
    }),
    actions: {
        async loadEmails(userEmail: string, field: string) {
            try {
                const data = { term: userEmail, field: field };
                const fetchedEmails = await apiHandleEmails('/ByField', {method: 'POST', data },)
                this.emails = fetchedEmails
            } catch (err) {
            this.error = err instanceof Error ? err.message : 'Error loading emails'
            console.error('Error loading emails:', this.error)
            }
        },
    },
});