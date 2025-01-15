import { defineStore } from 'pinia';
import type { EmailDocument } from '@/types';
import { useZincService } from '@/services/utils'

const { apiHandleEmails } = useZincService()

export const useEmailStore = defineStore('email', {
    state: () => ({
        emails: [] as EmailDocument[],
        error: null as string | null,
        totalResults: 0,
        itemsPerPage: 0, 
        searchQuery: '', 
        searchField: 'body', // por defecto
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
        async loadEmails() {
            try {
                const termOfSearch = this.searchQuery === '' ? '' : this.searchQuery;
                const maxResults = this.itemsPerPage + 1;
                const searchType = termOfSearch === '' ? 'match_all' : 'match'
                const data = { term: termOfSearch, field: this.searchField, max_results: 30 * maxResults, search_type: searchType };
                const response = await apiHandleEmails('/ByField', {method: 'POST', data },)

                this.emails = response.hits.hits.map((hit: any) => hit._source);
                this.totalResults = response.hits.total.value
            } catch (err) {
                this.error = err instanceof Error ? err.message : 'Error loading emails'
                console.error('Error loading emails:', this.error)
            }
        },
    },
});