<script lang="ts">
    import FileUpload from "./FileUpload.svelte";
</script>

<div class="animate-fade-in-up">
    <!-- Hero Section -->
    <div class="relative overflow-hidden pt-12 pb-16 sm:pb-24">
        <div
            class="relative mx-auto max-w-7xl px-4 sm:px-6 lg:px-8 text-center"
        >
            <h1
                class="text-4xl font-extrabold tracking-tight text-slate-900 sm:text-5xl md:text-6xl"
            >
                <span class="block">Stop guessing where</span>
                <span class="block text-indigo-600">your money goes.</span>
            </h1>
            <p
                class="mx-auto mt-3 max-w-md text-base text-slate-500 sm:text-lg md:mt-5 md:max-w-3xl md:text-xl"
            >
                Spendsense uses <strong>local AI</strong> to analyze your bank statements
                instantly. Private, fast, and smarter than your banking app.
            </p>

            <div
                class="mt-10 max-w-sm mx-auto sm:max-w-none sm:flex sm:justify-center gap-4"
            >
                <div
                    class="w-full max-w-md bg-white p-2 rounded-xl shadow-xl shadow-indigo-100 border border-slate-100"
                >
                    <FileUpload />
                </div>
            </div>
            <p class="mt-3 text-sm text-slate-400">
                Trusted by privacy-conscious users. Runs 100% offline.
            </p>
        </div>
    </div>

    <!-- Feature Grid -->
    <div class="py-12 bg-white border-y border-slate-100">
        <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
            <div class="grid grid-cols-1 gap-8 sm:grid-cols-2 lg:grid-cols-3">
                <div class="pt-6">
                    <div
                        class="flow-root bg-slate-50 rounded-2xl px-6 pb-8 h-full border border-slate-100 hover:border-indigo-200 transition-colors"
                    >
                        <div class="-mt-6">
                            <div>
                                <span
                                    class="inline-flex items-center justify-center rounded-xl bg-indigo-500 p-3 shadow-lg shadow-indigo-200"
                                >
                                    <svg
                                        class="h-6 w-6 text-white"
                                        fill="none"
                                        viewBox="0 0 24 24"
                                        stroke="currentColor"
                                    >
                                        <path
                                            stroke-linecap="round"
                                            stroke-linejoin="round"
                                            stroke-width="2"
                                            d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"
                                        />
                                    </svg>
                                </span>
                            </div>
                            <h3
                                class="mt-8 text-lg font-bold tracking-tight text-slate-900"
                            >
                                Privacy First
                            </h3>
                            <p class="mt-5 text-base text-slate-500">
                                Your financial data never leaves your device. We
                                use <strong>Ollama</strong> to run AI models locally
                                on your machine.
                            </p>
                        </div>
                    </div>
                </div>

                <div class="pt-6">
                    <div
                        class="flow-root bg-slate-50 rounded-2xl px-6 pb-8 h-full border border-slate-100 hover:border-indigo-200 transition-colors"
                    >
                        <div class="-mt-6">
                            <div>
                                <span
                                    class="inline-flex items-center justify-center rounded-xl bg-indigo-500 p-3 shadow-lg shadow-indigo-200"
                                >
                                    <svg
                                        class="h-6 w-6 text-white"
                                        fill="none"
                                        viewBox="0 0 24 24"
                                        stroke="currentColor"
                                    >
                                        <path
                                            stroke-linecap="round"
                                            stroke-linejoin="round"
                                            stroke-width="2"
                                            d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"
                                        />
                                    </svg>
                                </span>
                            </div>
                            <h3
                                class="mt-8 text-lg font-bold tracking-tight text-slate-900"
                            >
                                Smart Categorization
                            </h3>
                            <p class="mt-5 text-base text-slate-500">
                                Automatically tags your expenses (Food, Rent,
                                Subs). See exactly where your paycheck is going
                                every month.
                            </p>
                        </div>
                    </div>
                </div>

                <div class="pt-6">
                    <div
                        class="flow-root bg-slate-50 rounded-2xl px-6 pb-8 h-full border border-slate-100 hover:border-indigo-200 transition-colors"
                    >
                        <div class="-mt-6">
                            <div>
                                <span
                                    class="inline-flex items-center justify-center rounded-xl bg-indigo-500 p-3 shadow-lg shadow-indigo-200"
                                >
                                    <svg
                                        class="h-6 w-6 text-white"
                                        fill="none"
                                        viewBox="0 0 24 24"
                                        stroke="currentColor"
                                    >
                                        <path
                                            stroke-linecap="round"
                                            stroke-linejoin="round"
                                            stroke-width="2"
                                            d="M19.428 15.428a2 2 0 00-1.022-.547l-2.384-.477a6 6 0 00-3.86.517l-.318.158a6 6 0 01-3.86.517L6.05 15.21a2 2 0 00-1.806.547M8 4h8l-1 1v5.172a2 2 0 00.586 1.414l5 5c1.26 1.26.367 3.414-1.415 3.414H4.828c-1.782 0-2.674-2.154-1.414-3.414l5-5A2 2 0 009 10.172V5L8 4z"
                                        />
                                    </svg>
                                </span>
                            </div>
                            <h3
                                class="mt-8 text-lg font-bold tracking-tight text-slate-900"
                            >
                                AI Financial Coach
                            </h3>
                            <p class="mt-5 text-base text-slate-500">
                                Chat with your finances. Ask "How can I save
                                more?" or "Identify waste" and get instant,
                                personalized advice.
                            </p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>

    <!-- How It Works -->
    <div class="py-16 max-w-3xl mx-auto px-4 sm:px-6 lg:px-8">
        <h2 class="text-2xl font-bold text-center text-slate-900 mb-12">
            How it works
        </h2>
        <div class="relative">
            <!-- Line -->
            <div
                class="absolute left-1/2 top-0 bottom-0 w-0.5 bg-slate-200 transform -translate-x-1/2 hidden md:block"
            ></div>

            <div class="space-y-12">
                <!-- Step 1 -->
                <div
                    class="relative flex items-center justify-between md:justify-center group"
                >
                    <div class="flex flex-col md:flex-row items-center w-full">
                        <div
                            class="md:w-1/2 md:pr-8 md:text-right order-2 md:order-1 mt-4 md:mt-0"
                        >
                            <h3 class="font-bold text-lg text-slate-900">
                                1. Export Statement
                            </h3>
                            <p class="text-slate-500">
                                Download your transaction history (CSV) from
                                your bank's website.
                            </p>
                        </div>
                        <div
                            class="absolute left-1/2 transform -translate-x-1/2 flex items-center justify-center w-8 h-8 rounded-full bg-white border-2 border-indigo-500 z-10 order-1 md:order-2"
                        >
                            <span class="text-indigo-600 font-bold text-sm"
                                >1</span
                            >
                        </div>
                        <div
                            class="md:w-1/2 md:pl-8 order-3 md:order-3 hidden md:block"
                        ></div>
                    </div>
                </div>

                <!-- Step 2 -->
                <div
                    class="relative flex items-center justify-between md:justify-center group"
                >
                    <div class="flex flex-col md:flex-row items-center w-full">
                        <div
                            class="md:w-1/2 md:pr-8 md:text-right order-3 md:order-1 hidden md:block"
                        ></div>
                        <div
                            class="absolute left-1/2 transform -translate-x-1/2 flex items-center justify-center w-8 h-8 rounded-full bg-white border-2 border-indigo-500 z-10 order-1 md:order-2"
                        >
                            <span class="text-indigo-600 font-bold text-sm"
                                >2</span
                            >
                        </div>
                        <div
                            class="md:w-1/2 md:pl-8 order-2 md:order-3 mt-4 md:mt-0 text-center md:text-left"
                        >
                            <h3 class="font-bold text-lg text-slate-900">
                                2. Upload & Analyze
                            </h3>
                            <p class="text-slate-500">
                                Drag and drop the file here. We parse it
                                instantly.
                            </p>
                        </div>
                    </div>
                </div>

                <!-- Step 3 -->
                <div
                    class="relative flex items-center justify-between md:justify-center group"
                >
                    <div class="flex flex-col md:flex-row items-center w-full">
                        <div
                            class="md:w-1/2 md:pr-8 md:text-right order-2 md:order-1 mt-4 md:mt-0"
                        >
                            <h3 class="font-bold text-lg text-slate-900">
                                3. Get Insights
                            </h3>
                            <p class="text-slate-500">
                                See your dashboard and chat with the AI to find
                                savings.
                            </p>
                        </div>
                        <div
                            class="absolute left-1/2 transform -translate-x-1/2 flex items-center justify-center w-8 h-8 rounded-full bg-white border-2 border-indigo-500 z-10 order-1 md:order-2"
                        >
                            <span class="text-indigo-600 font-bold text-sm"
                                >3</span
                            >
                        </div>
                        <div
                            class="md:w-1/2 md:pl-8 order-3 md:order-3 hidden md:block"
                        ></div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>
