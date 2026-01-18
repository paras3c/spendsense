import { writable } from 'svelte/store';

export const dashboard = writable<any>(null);
export const expenses = writable<any[]>([]);
export const insights = writable<any[]>([]);
export const aiMode = writable<'polite' | 'savage'>('polite');
export const loading = writable<boolean>(false);
export const persona = writable<any>(null);
