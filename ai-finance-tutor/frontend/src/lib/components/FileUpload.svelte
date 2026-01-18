<script lang="ts">
    import { uploadCSV, getSampleData } from '../api';
    import { loading } from '../stores';
</script>

<div class="flex flex-col gap-6 p-8 max-w-md mx-auto bg-white rounded-xl shadow-sm border border-gray-100">
    <h1 class="text-3xl font-bold text-center text-gray-800">Finance Tutor</h1>
    
    <label class="border-2 border-dashed border-gray-300 p-8 rounded-lg cursor-pointer hover:bg-gray-50 flex flex-col items-center justify-center transition-colors">
        <input
            type="file"
            accept=".csv"
            on:change={(e) => e.target.files && uploadCSV(e.target.files[0])}
            disabled={$loading}
            class="hidden"
        />
        <svg xmlns="http://www.w3.org/2000/svg" class="h-10 w-10 text-gray-400 mb-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12" />
        </svg>
        <p class="text-center text-gray-600">Upload your CSV</p>
    </label>

    <div class="relative">
        <div class="absolute inset-0 flex items-center">
            <div class="w-full border-t border-gray-200"></div>
        </div>
        <div class="relative flex justify-center text-sm">
            <span class="px-2 bg-white text-gray-500">Or</span>
        </div>
    </div>

    <button
        on:click={getSampleData}
        disabled={$loading}
        class="bg-blue-600 text-white py-2 px-4 rounded-lg hover:bg-blue-700 disabled:opacity-50 font-medium transition-colors"
    >
        {$loading ? 'Loading...' : 'Use Sample Data'}
    </button>
</div>
