import { get } from 'svelte/store';
import { dashboard, expenses, insights, loading } from './stores';

const API_URL = 'http://localhost:8000';

export async function uploadCSV(file: File) {
    loading.set(true);
    const formData = new FormData();
    formData.append('file', file);

    try {
        const response = await fetch(`${API_URL}/upload`, {
            method: 'POST',
            body: formData,
        });
        if (!response.ok) throw new Error('Upload failed');
        const data = await response.json();
        updateStore(data);
    } catch (e) {
        console.error(e);
        alert('Upload failed: ' + e.message);
    } finally {
        loading.set(false);
    }
}

export async function getSampleData() {
    loading.set(true);
    try {
        const response = await fetch(`${API_URL}/sample-data`);
        if (!response.ok) throw new Error('Failed to get sample data');
        const data = await response.json();
        updateStore(data);
    } catch (e) {
        console.error(e);
        alert('Failed to load sample data');
    } finally {
        loading.set(false);
    }
}

export async function explainInsight(insight: any, followUp?: string, style: 'polite' | 'savage' = 'polite', action: 'explain' | 'draft_cancel' = 'explain') {
    const response = await fetch(`${API_URL}/explain-insight`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
            insight,
            follow_up: followUp,
            style: style,
            action: action
        }),
    });
    if (!response.ok) throw new Error('Failed to get explanation');
    return await response.json();
}

export async function generatePersona() {
    const response = await fetch(`${API_URL}/generate-persona`, {
        method: 'POST',
    });
    if (!response.ok) throw new Error('Failed to generate persona');
    return await response.json();
}

function updateStore(data: any) {
    dashboard.set(data);
    expenses.set(data.expenses);
    insights.set(data.insights);
}
