<script lang="ts">
    import * as Select from "$lib/components/ui/select";
    import { Separator } from "$lib/components/ui/separator/index.js";
    import { Button } from "$lib/components/ui/button/index.js";
    import * as ToggleGroup from "$lib/components/ui/toggle-group/index.js";
    import LineChart from "$lib/components/LineChart.svelte";

    // USER STATS
    let { data } = $props()
    let { graphDataPoints, user, watchlistGraphDataPoints } = data

    let pfVal: number = $state(graphDataPoints.at(-1)?.resalePrice)!
    let mortgageVal: number = $state(user.mortgage)
    let pfValCommas: string = $derived(numberWithCommas(pfVal))
    let mortgageValCommas: string = $derived(numberWithCommas(mortgageVal))
    
    let windowStr: string = $state("3650")
    let window: number = $derived(Number(windowStr))

    let pfDelta: number = $derived(graphDataPoints.at(-1)?.resalePrice! - (graphDataPoints.at(-window)?.resalePrice || graphDataPoints[0].resalePrice))
    let pfDeltaCommas: string = $derived(numberWithCommas(pfDelta < 0 ? -pfDelta : pfDelta))
    let pfPercentDelta: number = $derived(pfDelta / pfVal * 100)

    let pfValueGraphData: DataSet[] = $derived([
        {
            label: "",
			data: graphDataPoints.slice(graphDataPoints.length - 1 - window).map(point => point.resalePrice),
			xAxis: graphDataPoints.slice(graphDataPoints.length - 1 - window).map(point => point.date),
			borderColor: "#4BAAC8",
			backgroundColor: "#4BAAC8"
        }
    ])

    let watchlistGraphData: DataSet[][] = $derived(watchlistGraphDataPoints.map((assetGraphDataPoints) => [
        {
            label: "",
            data: assetGraphDataPoints.slice(assetGraphDataPoints.length - 1 - window).map((point) => point.resalePrice),
            xAxis: assetGraphDataPoints.slice(assetGraphDataPoints.length - 1 - window).map((point) => point.date),
            borderColor: "#FF8C00",
            backgroundColor: "#FF8C00",
        },
    ]))

    function numberWithCommas(n: number): string {
        return n.toFixed(2).replace(/\B(?=(\d{3})+(?!\d))/g, ",");
    } 


    let watchListedAssets: Asset[] = $derived([
        {
            town: "QUEENSTOWN",
            flatType: "4 ROOM",
            graph: watchlistGraphData[0],
            resalePrice: 502234.32,
            percentageChange: 3.24
        },
        {
            town: "CLEMENTI",
            flatType: "4 ROOM",
            graph: watchlistGraphData[1],
            resalePrice: 632042.30,
            percentageChange: -2.13
        },
        {
            town: "BISHAN",
            flatType: "4 ROOM",
            graph: watchlistGraphData[2],
            resalePrice: 435022.20,
            percentageChange: 0.12
        },
    ])

    function deAllCaps(str: string): string {
        str = str.toLowerCase()
        let words = str.split(" ")
        for (let i = 0; i < words.length; i++) {
           words[i] = words[i].charAt(0).toUpperCase() + words[i].substring(1)
        }
        return words.join(" ")
    }


</script>

<div class="p-5">
    <!-- PERSONAL STATS -->
    <div class="mb-10">
        <!-- Value of Properties Owned -->
         <div class="pl-4 mb-3">
             <h1 class="text-3xl">${pfValCommas}</h1>
             <p class="text-md {pfDelta > 0 ? " text-green-400" : " text-red-400"}">{@html pfDelta > 0 ? "&#8599;" : "&#x2198;"} ${pfDeltaCommas} ({pfPercentDelta.toFixed(2)}%)</p>
         </div>
         
         <!-- Graph -->
         <LineChart
         {...{
             stats: pfValueGraphData,
             label: "Value (SGD)"
         }} />	
         
        <ToggleGroup.Root class="mb-6" type="single" bind:value={windowStr}>
            <ToggleGroup.Item value=30>
                1M
            </ToggleGroup.Item>
            <ToggleGroup.Item value=90>
                3M
            </ToggleGroup.Item>
            <ToggleGroup.Item value=365>
                1Y
            </ToggleGroup.Item>
            <ToggleGroup.Item value=1825>
                5Y
            </ToggleGroup.Item>
            <ToggleGroup.Item value=3650>
                10Y
            </ToggleGroup.Item>
            <ToggleGroup.Item value="0{graphDataPoints.length-1}">
                All
            </ToggleGroup.Item>
        </ToggleGroup.Root>

        <div class="px-3">
            <div class="flex w-full justify-between hover:bg-zinc-900 p-2 px-4 rounded-md">
                <div>Assets</div>
                <div>${pfValCommas} &#10095;</div>
            </div>
            <div class="flex w-full justify-between hover:bg-zinc-900 p-2 px-4 rounded-md">
                <div>Mortgage</div>
                <div>${mortgageValCommas} &#10095;</div>
            </div>
        </div>
    </div>

    <Separator class="my-8" />
    
    <!-- WATCHLIST -->
     <div class="mb-10">
         <div class="flex justify-between mb-4">
            <h1 class="text-xl pl-4">Watchlist</h1>
    
            <Select.Root>
                <Select.Trigger class="w-[180px]">
                  <Select.Value placeholder="Filter" />
                </Select.Trigger>
                <Select.Content>
                  <Select.Item value="watchlist">Watchlist</Select.Item>
                  <Select.Item value="top">Top Assets</Select.Item>
                  <Select.Item value="trending">Trending</Select.Item>
                </Select.Content>
            </Select.Root>
        </div>

        <!-- Current price, Percentage Change -->
        <div class="mb-5">
            {#each watchListedAssets as asset}
            <a href="/stats/{asset.town}+{asset.flatType}">
                <div class="flex justify-between hover:bg-zinc-900 p-2 px-4 rounded-md">
                    <div>
                        <p>{deAllCaps(asset.town)}</p>
                        <p class="text-sm text-gray-500">{asset.flatType}</p>
                    </div>

                    <div class="w-24 h-20">
                        <LineChart
                        {...{
                            stats: asset.graph,
                            label: "Value (SGD)"
                        }} />	
                    </div>
    
                    <div class="flex flex-col items-end">
                        <p>${asset.resalePrice.toFixed(2)}</p>
                        <p class="text-sm p-1 text-end w-[70px] rounded-lg {asset.percentageChange > 0 ? "bg-green-800 text-green-300" : "bg-red-800 text-red-300"}">{@html asset.percentageChange > 0 ? "&#8599;" : "&#x2198;"} {asset.percentageChange.toFixed(2)}%</p>
                    </div>
                </div>
            </a>
            {/each}
        </div>

        <div class="flex justify-center">
            <Button class="w-3/4 rounded-3xl bg-secondary hover:bg-gray-600 text-primary">Manage</Button>
        </div>

     </div>
     
     <Separator class="my-8" />

     <div>
        <div class="pl-4 mb-4">
            <h1 class="text-xl">Leaderboards</h1>
            <p class="text-sm text-gray-500">Based on the past 30 days</p>
        </div>
         <!-- Top gainers, Top losers, Most listings -->
         <div class="border-2 w-full h-60 flex justify-center items-center mb-5">LEADERBOARDS PLACEHOLDER</div>

        <div class="flex justify-center">
            <Button class="w-3/4 rounded-3xl bg-secondary hover:bg-gray-600 text-primary">Explore Leaderboards</Button>
        </div>
     </div>
</div>