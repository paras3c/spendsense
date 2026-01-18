<script lang="ts">
    import { scale } from "svelte/transition";

    interface Persona {
        archetype: string;
        emoji: string;
        description: string;
        savage_quote: string;
    }

    export let data: Persona;
</script>

<div
    class="relative w-full mx-auto overflow-hidden bg-white border-[3px] border-black p-8 shadow-[12px_12px_0px_0px_rgba(0,0,0,1)] transition-transform hover:-translate-y-1"
    in:scale
>
    <!-- Editorial "Photo Frame" Layout -->
    <div class="flex flex-col items-center text-center">
        <!-- Icon/Avatar Placeholder -->
        <div
            class="w-32 h-32 border-2 border-black rounded-full flex items-center justify-center mb-6 bg-gray-50"
        >
            <span class="text-6xl filter grayscale contrast-125"
                >{data.emoji}</span
            >
        </div>

        <div class="inline-block border-b-2 border-black pb-1 mb-4">
            <span
                class="text-xs font-bold uppercase tracking-[0.2em] text-gray-500"
                >Financial Archetype</span
            >
        </div>

        <h2
            class="text-5xl font-serif font-black mb-6 text-black leading-none tracking-tight"
        >
            {data.archetype}
        </h2>

        <p
            class="text-lg font-serif text-gray-800 leading-relaxed max-w-sm mx-auto mb-8 border-l-4 border-black pl-4 text-left"
        >
            {data.description}
        </p>

        <!-- Savage Quote Section -->
        <div class="w-full bg-black text-white p-6 relative">
            <div
                class="absolute -top-3 left-1/2 transform -translate-x-1/2 bg-white px-2 border border-black"
            >
                <span
                    class="text-[10px] font-bold uppercase tracking-widest text-black"
                    >Analysis</span
                >
            </div>
            <p class="font-serif italic text-xl leading-relaxed">
                "{data.savage_quote}"
            </p>
        </div>

        <div class="mt-8 flex items-center gap-2">
            <div class="h-px bg-black w-12"></div>
            <p
                class="text-[10px] font-mono tracking-widest uppercase text-black"
            >
                AI Finance Tutor
            </p>
            <div class="h-px bg-black w-12"></div>
        </div>
    </div>
</div>
