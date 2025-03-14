<script lang="ts">
    import * as Select from "$lib/components/ui/select";
    import { Separator } from "$lib/components/ui/separator/index.js";
    import { Button } from "$lib/components/ui/button/index.js";

    // USER STATS
    let assets: Asset[] = []
    let pfVal: number = $state(621433.32)
    let pfValCommas: string = $derived(numberWithCommas(pfVal))

    let pfDelta: number = $state(-73320.12)
    let pfDeltaCommas: string = $derived(numberWithCommas(pfDelta < 0 ? -pfDelta : pfDelta))
    let pfPercentDelta: number = $derived(pfDelta / pfVal * 100)

    function numberWithCommas(n: number): string {
        return n.toFixed(2).replace(/\B(?=(\d{3})+(?!\d))/g, ",");
    } 


    let watchListedAssets: Asset[] = [
        {
            town: "QUEENSTOWN",
            flatType: "4 ROOM",
            resalePrice: 502234.32,
            percentageChange: 3.24
        },
        {
            town: "CLEMENTI",
            flatType: "4 ROOM",
            resalePrice: 632042.30,
            percentageChange: -2.13
        },
        {
            town: "BISHAN",
            flatType: "4 ROOM",
            resalePrice: 435022.20,
            percentageChange: 0.12
        },
    ]

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
    <div class="mb-10 pl-4">
        <!-- Value of Properties Owned -->
         <h1 class="text-3xl">${pfValCommas}</h1>
         <p class="text-md {pfDelta > 0 ? " text-green-400" : " text-red-400"}">{@html pfDelta > 0 ? "&#8599;" : "&#x2198;"} ${pfDeltaCommas} ({pfPercentDelta.toFixed(2)}%)</p>
         <!-- Graph -->
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
         {#each watchListedAssets as asset}
         <a href="/stats/{asset.town}+{asset.flatType}">
             <div class="flex justify-between hover:bg-zinc-900 p-2 px-4 rounded-md">
                 <div>
                     <p>{deAllCaps(asset.town)}</p>
                     <p class="text-sm text-gray-500">{asset.flatType}</p>
                 </div>
 
                 <div class="flex flex-col items-end">
                     <p>${asset.resalePrice.toFixed(2)}</p>
                     <p class="text-sm p-1 text-end w-[70px] rounded-lg {asset.percentageChange > 0 ? "bg-green-800 text-green-300" : "bg-red-800 text-red-300"}">{@html asset.percentageChange > 0 ? "&#8599;" : "&#x2198;"} {asset.percentageChange.toFixed(2)}%</p>
                 </div>
             </div>
         </a>
         {/each}

         <Button class="w-3/4 mx-16 mt-5 rounded-3xl bg-secondary hover:bg-gray-600 text-primary">Manage</Button>

     </div>
     
     <Separator class="my-8" />

     <div>
        <div class="pl-4">
            <h1 class="text-xl">Leaderboards</h1>
            <p class="text-sm text-gray-500">Based on the past 30 days</p>
        </div>
         <!-- Top gainers, Top losers, Most listings -->

         <Button class="w-3/4 mx-16 mt-5 rounded-3xl bg-secondary hover:bg-gray-600 text-primary">Explore Leaderboards</Button>
     </div>
</div>