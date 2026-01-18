<script lang="ts">
    import {
        dashboard,
        expenses,
        insights,
        aiMode,
        persona,
        loading,
    } from "$lib/stores";
    import LandingPage from "$lib/components/LandingPage.svelte";
    import InsightCard from "$lib/components/InsightCard.svelte";
    import SpendingChart from "$lib/components/SpendingChart.svelte";
    import PersonaCard from "$lib/components/PersonaCard.svelte";
    import { generatePersona } from "$lib/api";
    import { fade } from "svelte/transition";
    import { onMount } from "svelte";

    function reset() {
        dashboard.set(null);
        expenses.set([]);
        insights.set([]);
        persona.set(null);
        // Clear history state if present
        if (history.state?.view === "dashboard") {
            history.back();
        }
    }

    // Handle browser back button
    onMount(() => {
        const handlePopState = (event: PopStateEvent) => {
            if (!event.state || event.state.view !== "dashboard") {
                // If we go back and state is null (landing), reset stores
                dashboard.set(null);
                expenses.set([]);
                insights.set([]);
                persona.set(null);
            }
        };

        window.addEventListener("popstate", handlePopState);
        return () => window.removeEventListener("popstate", handlePopState);
    });

    // Subscribe to dashboard changes to push state
    $: if (
        $dashboard &&
        (!history.state || history.state.view !== "dashboard")
    ) {
        history.pushState({ view: "dashboard" }, "");
    }

    function toggleMode() {
        aiMode.update((m) => (m === "polite" ? "savage" : "polite"));
    }

    async function handlePersona() {
        loading.set(true);
        try {
            const res = await generatePersona();
            persona.set(res);
        } catch (e) {
            console.error(e);
            alert("Could not generate persona");
        }
        loading.set(false);
    }
</script>

<div
    class="min-h-screen font-sans text-black bg-white transition-colors duration-500"
>
    <!-- Top Navigation -->
    <nav class="editorial-border-b bg-white sticky top-0 z-10">
        <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
            <div class="flex justify-between h-20 items-center">
                <div class="flex items-center gap-4">
                    <div
                        class="p-2 border-2 border-black rounded-none transition-colors hover:bg-black hover:text-white"
                    >
                        <svg
                            xmlns="http://www.w3.org/2000/svg"
                            class="h-6 w-6"
                            fill="none"
                            viewBox="0 0 24 24"
                            stroke="currentColor"
                        >
                            <path
                                stroke-linecap="round"
                                stroke-linejoin="round"
                                stroke-width="2"
                                d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
                            />
                        </svg>
                    </div>
                    <span
                        class="font-serif font-bold text-3xl tracking-tight text-black"
                        >Spendsense.</span
                    >
                </div>
                {#if $dashboard}
                    <div class="flex items-center gap-4">
                        <!-- Mode Toggle -->
                        <button
                            on:click={toggleMode}
                            class="editorial-btn-outline text-xs"
                        >
                            {#if $aiMode === "savage"}
                                SAVAGE MODE: ON
                            {:else}
                                SAVAGE MODE: OFF
                            {/if}
                        </button>

                        <button on:click={reset} class="editorial-btn text-xs">
                            RESET
                        </button>
                    </div>
                {/if}
            </div>
        </div>
    </nav>

    <main class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
        {#if !$dashboard}
            <LandingPage />
        {:else}
            <!-- Dashboard Content -->
            <div class="space-y-12 animate-fade-in-up">
                <!-- Header Section -->
                <div class="border-b-2 border-black pb-8">
                    <h1 class="text-6xl font-serif font-bold mb-2">
                        Financial Overview
                    </h1>
                    <p class="text-xl font-serif italic text-gray-600">
                        Your spending habits, analyzed.
                    </p>
                </div>

                <!-- Stats Grid -->
                <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
                    <!-- Stat 1 -->
                    <div class="editorial-card relative group">
                        <div
                            class="absolute top-4 right-4 text-black opacity-10 group-hover:opacity-20 transition-opacity"
                        >
                            <svg
                                xmlns="http://www.w3.org/2000/svg"
                                class="h-20 w-20"
                                fill="none"
                                viewBox="0 0 24 24"
                                stroke="currentColor"
                            >
                                <path
                                    stroke-linecap="round"
                                    stroke-linejoin="round"
                                    stroke-width="1"
                                    d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
                                />
                            </svg>
                        </div>
                        <p
                            class="font-serif font-bold text-lg mb-4 border-b border-black pb-2 inline-block"
                        >
                            TOTAL SPEND
                        </p>
                        <div class="flex items-baseline gap-1 mt-4">
                            <span class="text-2xl font-serif">₹</span>
                            <span
                                class="text-6xl font-bold font-serif tracking-tighter"
                            >
                                {$dashboard.total_expenses.toLocaleString(
                                    "en-IN",
                                )}
                            </span>
                        </div>
                    </div>

                    <!-- Stat 2 -->
                    <div class="editorial-card group">
                        <p
                            class="font-serif font-bold text-lg mb-4 border-b border-black pb-2 inline-block"
                        >
                            DAILY AVERAGE
                        </p>
                        <div class="flex items-baseline gap-1 mt-4">
                            <span class="text-2xl font-serif">₹</span>
                            <span
                                class="text-6xl font-bold font-serif tracking-tighter"
                            >
                                {$dashboard.average_daily.toLocaleString(
                                    "en-IN",
                                )}
                            </span>
                        </div>
                    </div>

                    <!-- Stat 3: Confidence -->
                    <div class="editorial-card group relative overflow-hidden">
                        <p
                            class="font-serif font-bold text-lg mb-4 border-b border-black pb-2 inline-block"
                        >
                            CONFIDENCE SCORE
                        </p>
                        <div
                            class="flex items-baseline gap-2 mt-4 relative z-10"
                        >
                            <span
                                class="text-7xl font-black font-serif tracking-tighter"
                            >
                                {$dashboard.confidence_score || 0}
                            </span>
                            <span
                                class="text-sm font-bold uppercase tracking-widest text-gray-500"
                                >/100</span
                            >
                        </div>
                        <div
                            class="mt-4 font-serif text-lg italic border-t border-gray-200 pt-2"
                        >
                            {($dashboard.confidence_score || 0) >= 80
                                ? "Status: Excellent"
                                : ($dashboard.confidence_score || 0) >= 50
                                  ? "Status: Optimization Required"
                                  : "Status: Critical Action Needed"}
                        </div>
                    </div>
                </div>

                <!-- Persona Section -->
                <div
                    class="editorial-card cursor-pointer hover:bg-gray-50 transition-colors"
                    on:click={handlePersona}
                >
                    <div
                        class="flex flex-col items-center justify-center p-8 text-center border-2 border-dashed border-gray-300 hover:border-black transition-colors"
                    >
                        {#if !$persona}
                            <h3 class="text-3xl font-serif font-bold mb-2">
                                Generate Financial Persona
                            </h3>
                            <p class="font-serif italic text-gray-500">
                                Discover your spending archetype with AI.
                            </p>
                            <button class="editorial-btn mt-6"
                                >Analyze Now</button
                            >
                        {:else}
                            <h3 class="text-4xl font-serif font-bold mb-2">
                                {$persona.archetype}
                            </h3>
                            <button
                                class="editorial-btn-outline mt-4"
                                on:click|stopPropagation={() =>
                                    persona.set(null)}>Reset Persona</button
                            >
                            যাবে
                        {/if}
                    </div>
                </div>

                <!-- Full Persona Overlay -->
                {#if $persona}
                    <!-- svelte-ignore a11y-click-events-have-key-events -->
                    <div
                        class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-white/90 backdrop-blur-sm transition-all"
                        on:click={() => persona.set(null)}
                        transition:fade={{ duration: 200 }}
                    >
                        <div
                            on:click|stopPropagation
                            class="editorial-card shadow-2xl max-w-2xl w-full"
                        >
                            <PersonaCard data={$persona} />
                        </div>
                    </div>
                {/if}

                <!-- Main Content Grid -->
                <div
                    class="grid grid-cols-1 lg:grid-cols-12 gap-12 pt-8 border-t-2 border-black"
                >
                    <!-- Left Column: Insights -->
                    <div class="lg:col-span-7 space-y-8">
                        <div
                            class="flex items-center justify-between border-b border-black pb-4"
                        >
                            <h2 class="text-4xl font-serif font-bold">
                                AI Intelligence
                            </h2>
                            <span
                                class="bg-black text-white px-3 py-1 font-bold text-xs uppercase tracking-widest"
                            >
                                {$insights.length} Insights
                            </span>
                        </div>
                        <div class="space-y-6">
                            {#each $insights as insight}
                                {#if insight.type !== "category_breakdown"}
                                    <InsightCard {insight} />
                                {/if}
                            {/each}
                        </div>
                    </div>

                    <!-- Right Column: Chart -->
                    <div class="lg:col-span-5 space-y-8">
                        <h2
                            class="text-4xl font-serif font-bold border-b border-black pb-4"
                        >
                            Expenditure Map
                        </h2>
                        <div
                            class="editorial-card min-h-[400px] flex items-center justify-center"
                        >
                            {#if $insights.find((i) => i.type === "category_breakdown")}
                                <SpendingChart
                                    breakdown={$insights.find(
                                        (i) => i.type === "category_breakdown",
                                    ).breakdown}
                                />
                            {:else}
                                <div class="text-center py-12">
                                    <p
                                        class="font-serif italic text-gray-400 text-xl"
                                    >
                                        No category data available
                                    </p>
                                </div>
                            {/if}
                        </div>
                    </div>
                </div>
            </div>
        {/if}
    </main>
</div>

<style>
    .animate-fade-in-up {
        animation: fadeInUp 0.8s cubic-bezier(0.16, 1, 0.3, 1) forwards;
    }
    @keyframes fadeInUp {
        from {
            opacity: 0;
            transform: translateY(40px);
        }
        to {
            opacity: 1;
            transform: translateY(0);
        }
    }
</style>
