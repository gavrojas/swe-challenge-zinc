import { defineStore } from 'pinia';
import type { EmailDocument } from '@/types/emails';
import { useZincService } from '@/services/zinc'

export const useEmailStore = defineStore('email', {
    state: () => ({
        emails: [] as EmailDocument[],
        error: null as string | null,
        folders: [
            'all_documents',
            'calendar',
            'deleted_items',
            'eol',
            'inbox',
            'nymex',
            'reuters',
            'tasks',
            'tradecounts',
            'associate_prc',
            'charges',
            'discussion_threads',
            'espeed',
            'kiodex',
            'offline',
            'sent',
            'to_do',
            'tss',
            'broker_client',
            'contacts',
            'emetra',
            'hs',
            'natsource',
            'origination',
            'sent_items',
            'tq',
        ],
    }),
    actions: {
        async loadEmails(userEmail: string) {
            const { emails, error, fetchEmails } = useZincService()
            try {
            await fetchEmails(userEmail)
            this.emails = emails.value
            this.error = error.value
            } catch (err) {
            this.error = err instanceof Error ? err.message : 'Error loading emails'
            console.error('Error loading emails:', this.error)
            }
        },
        },
});