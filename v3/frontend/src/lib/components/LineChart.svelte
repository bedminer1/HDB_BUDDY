<script lang="ts">
    import { onDestroy, onMount } from 'svelte';
    import { mode } from "mode-watcher"
    import { Chart, LineController, LineElement, PointElement, LinearScale, Title, CategoryScale, Tooltip, Legend } from 'chart.js';

    Chart.register(LineController, LineElement, PointElement, LinearScale, Title, CategoryScale, Tooltip, Legend);

    export let stats : DataSet[] 

    let chart: Chart | null = null
    let chartCanvas: HTMLCanvasElement
    $: Chart.defaults.color = $mode === "dark" ? 'rgb(250,255,255)' : 'rgb(0,0,0)'
    Chart.defaults.font.size = 14

    onMount(() => {
        if (stats && stats.length > 0 && chartCanvas) {
            createChart();
        }
    });

    $: {
        if (chart && $mode && stats && stats.length > 0 && chartCanvas) {
            chart.destroy();
            createChart();
        }
    }

    function createChart() {

        chart = new Chart(chartCanvas, {
            type: 'line',
            data: {
                labels: stats[0].xAxis, // Use the xAxis from the first dataset (assuming they are the same)
                datasets: stats.map(stat => ({
                    label: stat.label,
                    data: stat.data,
                    borderColor: stat.borderColor || 'rgba(75, 192, 192, 1)',
                    backgroundColor: stat.backgroundColor || 'rgba(75, 192, 192, 0.2)',
                    fill: true,
                    lineTension: 0.1,
                })),
            },
            options: {
                responsive: true,
                plugins: {
                    legend: {
                        display: false
                    },
                    tooltip: {
                        enabled: false
                    }
                },
                elements: {
                    point: {
                        radius: 0
                    }
                },
                scales: {
                    y: {
                        ticks: {
                            display: false,
                        }
                    },
                    x: {
                        ticks: {
                            display: false,
                        }
                    },
                },
            },
        });
    }

    onDestroy(() => {
        if (chart) chart.destroy();
    });
</script>

<canvas class="w-full rounded-md" bind:this={chartCanvas}></canvas>