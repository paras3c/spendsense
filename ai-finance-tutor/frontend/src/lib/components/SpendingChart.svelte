<script lang="ts">
    import { onMount, onDestroy } from "svelte";
    import {
        Chart,
        PieController,
        ArcElement,
        Tooltip,
        Legend,
    } from "chart.js";

    Chart.register(PieController, ArcElement, Tooltip, Legend);

    export let breakdown: Record<string, number> = {};

    let canvas: HTMLCanvasElement;
    let chart: Chart;

    onMount(() => {
        if (!canvas) return;

        chart = new Chart(canvas, {
            type: "doughnut",
            data: {
                labels: Object.keys(breakdown),
                datasets: [
                    {
                        data: Object.values(breakdown),
                        backgroundColor: [
                            "#000000", // Black
                            "#333333", // Dark Gray
                            "#666666", // Medium Gray
                            "#999999", // Light Gray
                            "#cccccc", // Lighter Gray
                            "#e5e5e5", // Very Light Gray
                        ],
                        hoverOffset: 10,
                        borderWidth: 2,
                        borderColor: "#ffffff",
                        hoverBorderColor: "#000000",
                    },
                ],
            },
            options: {
                responsive: true,
                maintainAspectRatio: false,
                cutout: "50%",
                plugins: {
                    legend: {
                        position: "bottom",
                        labels: {
                            padding: 20,
                            color: "#000000", // Black text
                            font: {
                                family: "'Playfair Display', serif",
                                size: 12,
                                weight: 700,
                            },
                            usePointStyle: true,
                            pointStyle: "rectRot", // Editorial diamond style
                        },
                    },
                    tooltip: {
                        backgroundColor: "#ffffff",
                        titleColor: "#000000",
                        bodyColor: "#333333",
                        titleFont: {
                            family: "'Playfair Display', serif",
                            size: 14,
                        },
                        bodyFont: { family: "'Lora', serif", size: 12 },
                        borderColor: "#000000",
                        borderWidth: 2,
                        padding: 12,
                        cornerRadius: 0, // Sharp corners
                        displayColors: true,
                        boxPadding: 4,
                    },
                },
                layout: {
                    padding: 20,
                },
                animation: {
                    animateScale: true,
                    animateRotate: true,
                },
            },
        });
    });

    onDestroy(() => {
        if (chart) chart.destroy();
    });
</script>

<div class="relative w-full h-80 mx-auto">
    <canvas bind:this={canvas}></canvas>
</div>
