import { defineStore } from 'pinia';
import { ref } from 'vue';
import type { EmailDocument } from '@/types';
import { useZincService } from '@/services/utils'

const { apiHandleEmails } = useZincService()

export const useEmailStore = defineStore('email',  () => {
    const emails = ref<EmailDocument[]>([]);
    const error = ref<string | null>(null);
    const totalResults = ref(0);
    const itemsPerPage = ref(0);
    const searchQuery = ref('');
    const searchField = ref('body'); // por defecto
    const folders = ref([
        'all_documents',
        'inbox',
        'sent',
        'sent_items',
        'deleted_items',
        'discussion_threads',
        'tasks',
        'calendar',
        'contacts',
    ]);

    const loadEmails = async () => {
        try {
            const termOfSearch = searchQuery.value === '' ? '' : searchQuery.value;
            const maxResults = itemsPerPage.value + 1;
            const searchType = termOfSearch === '' ? 'match_all' : 'match';
            const data = { term: termOfSearch, field: searchField.value, max_results: 30 * maxResults, search_type: searchType };
            const response = await apiHandleEmails('/ByField', { method: 'POST', data });

            emails.value = response.hits.hits.map((hit: any) => hit._source);
            totalResults.value = response.hits.total.value;
        } catch (err) {
            error.value = err instanceof Error ? err.message : 'Error loading emails';
            console.error('Error loading emails:', error.value);
        }
    };

    return { emails, error, totalResults, itemsPerPage, searchQuery, searchField, folders, loadEmails };
});