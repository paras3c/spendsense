<script lang="ts">
    import { explainInsight } from "$lib/api";
    import { aiMode } from "$lib/stores";

    export let insight: any;

    let showExplanation = false;
    let loadingExplain = false;

    // Editorial uses solid borders/text instead of colored glows
    const colors = {
        alert: "border-black bg-black text-white", // Dramatic for alerts
        warning: "border-black bg-white text-black", // Standard
        info: "border-gray-300 bg-white text-gray-500", // Subtle
    };

    const getIconPath = (type: string) => {
        switch (type) {
            case "subscription_waste":
                return "M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"; // Exclamation circle
            case "high_food":
                return "M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z"; // Shopping cart (approx for food)
            case "fixed_vs_variable":
                return "M3 6l3 1m0 0l-3 9a5.002 5.002 0 006.001 0M6 7l3 9M6 7l6-2m6 2l3-1m-3 1l-3 9a5.002 0 006.001 0M18 7l3 9m-3-9l-6-2m0-2v2m0 16V5m0 16H9m3 0h3"; // Scale/Balance
            default:
                return "M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"; // Info
        }
    };

    let messages: { role: "user" | "assistant"; content: string }[] = [];
    let userInput = "";

    const handleExplain = async (
        action: "explain" | "draft_cancel" = "explain",
    ) => {
        loadingExplain = true;
        showExplanation = true;

        // Add initial user message to chat UI for context
        if (action === "draft_cancel") {
            messages = [
                {
                    role: "user",
                    content: "Draft a cancellation email for this.",
                },
            ];
        } else {
            messages = []; // Clear for new explanation
        }

        try {
            // Pass the current mode (polite/savage) and action
            const res = await explainInsight(insight, "", $aiMode, action);
            messages = [
                ...messages,
                { role: "assistant", content: res.explanation },
            ];
        } catch (error) {
            console.error(error);
            messages = [
                {
                    role: "assistant",
                    content: "Sorry, I couldn't connect to the AI tutor.",
                },
            ];
        }
        loadingExplain = false;
    };

    const handleChat = async () => {
        if (!userInput.trim()) return;

        // Optimistic update
        const userMsg = userInput;
        messages = [...messages, { role: "user", content: userMsg }];
        userInput = "";
        loadingExplain = true;

        try {
            const res = await explainInsight(
                insight,
                userMsg,
                $aiMode,
                "explain",
            ); // Follow-ups continue in current mode
            messages = [
                ...messages,
                { role: "assistant", content: res.explanation },
            ];
        } catch (error) {
            messages = [
                ...messages,
                {
                    role: "assistant",
                    content: "Error: Could not fetch response.",
                },
            ];
        }
        loadingExplain = false;
    };
</script>

<!-- Clean Editorial Card -->
<div class="border-b border-black py-6 first:pt-0 last:border-0 group">
    <div class="flex items-start gap-4">
        <!-- Minimal Icon Box -->
        <div
            class="w-8 h-8 flex items-center justify-center border border-black rounded-none flex-shrink-0"
        >
            <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-5 w-5 text-black"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
            >
                <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="1.5"
                    d={getIconPath(insight.type)}
                />
            </svg>
        </div>

        <div class="flex-grow">
            <!-- Impact Context -->
            {#if insight.impact_context}
                <p class="text-xs font-serif italic text-gray-500 mb-1">
                    "{insight.impact_context}"
                </p>
            {/if}

            <h3
                class="font-serif font-bold text-xl leading-tight text-black mb-2"
            >
                {insight.message}
            </h3>

            <!-- Actionable Step -->
            {#if insight.actionable_step}
                <div class="border-l-2 border-black pl-3 my-3">
                    <p
                        class="text-xs font-bold uppercase tracking-widest text-black mb-1"
                    >
                        Recommendation
                    </p>
                    <p class="text-sm font-medium text-black">
                        {insight.actionable_step}
                    </p>
                </div>
            {/if}

            {#if insight.breakdown}
                <div class="mt-3 border border-dashed border-gray-400 p-3">
                    {#each Object.entries(insight.breakdown) as [cat, amt]}
                        <div
                            class="flex justify-between text-sm py-1 border-b border-gray-100 last:border-0"
                        >
                            <span class="text-gray-600 font-serif italic"
                                >{cat}</span
                            >
                            <span class="text-black font-bold font-mono"
                                >₹{amt}</span
                            >
                        </div>
                    {/each}
                </div>
            {:else}
                <p class="text-sm font-serif text-gray-600 mt-1">
                    Estimated Impact: <span class="text-black font-bold"
                        >₹{insight.monthly_cost.toLocaleString(
                            "en-IN",
                        )}/mo</span
                    >
                </p>
            {/if}

            {#if !showExplanation && !insight.breakdown}
                <div class="mt-4 flex gap-3">
                    <button
                        on:click={() => handleExplain("explain")}
                        disabled={loadingExplain}
                        class="editorial-btn-outline text-xs"
                    >
                        {loadingExplain
                            ? "Analyzing..."
                            : $aiMode === "savage"
                              ? "CRITICAL REVIEW"
                              : "AI ANALYSIS"}
                    </button>

                    <!-- Subscription Assassin Button -->
                    {#if insight.type === "subscription_waste"}
                        <button
                            on:click={() => handleExplain("draft_cancel")}
                            disabled={loadingExplain}
                            class="text-xs font-bold text-gray-500 hover:text-black underline transition-colors disabled:opacity-50"
                        >
                            Draft Cancellation
                        </button>
                    {/if}
                </div>
            {/if}
        </div>
    </div>

    <!-- Chat/Explanation Section -->
    {#if showExplanation}
        <div class="mt-5 pl-12">
            <div
                class="space-y-4 max-h-60 overflow-y-auto pr-2 custom-scrollbar bg-gray-50 p-4 border border-black"
            >
                {#each messages as msg}
                    <div
                        class={`flex ${msg.role === "user" ? "justify-end" : "justify-start"}`}
                    >
                        <div
                            class={`max-w-[90%] text-sm p-3 border ${
                                msg.role === "user"
                                    ? "bg-black text-white border-black"
                                    : "bg-white text-black border-black"
                            }`}
                        >
                            {msg.content}
                        </div>
                    </div>
                {/each}
                {#if loadingExplain}
                    <div class="text-xs font-mono text-gray-400">Typing...</div>
                {/if}
            </div>

            <div class="mt-4 flex gap-2">
                <input
                    type="text"
                    bind:value={userInput}
                    on:keydown={(e) => e.key === "Enter" && handleChat()}
                    placeholder="Ask a follow-up..."
                    class="block w-full border-b border-black text-sm p-2 focus:outline-none focus:border-b-2 bg-transparent placeholder-gray-400"
                />
                <button
                    on:click={handleChat}
                    disabled={!userInput || loadingExplain}
                    class="editorial-btn text-xs"
                >
                    SEND
                </button>
            </div>
        </div>
    {/if}
</div>

<style>
    .custom-scrollbar::-webkit-scrollbar {
        width: 4px;
    }
    .custom-scrollbar::-webkit-scrollbar-track {
        background: #f1f1f1;
    }
    .custom-scrollbar::-webkit-scrollbar-thumb {
        background: #000;
    }
</style>
