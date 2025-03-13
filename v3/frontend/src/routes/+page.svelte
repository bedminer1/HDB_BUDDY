<script lang="ts">
    import * as Select from "$lib/components/ui/select";

    // USER STATS
    let assets: Asset[] = []
    let portfolioValue: number = $state(0)


    let watchListedAssets: Asset[] = [
        {
            town: "QUEENSTOWN",
            flatType: "4 ROOM",
            resalePrice: 50234.32,
            percentageChange: 3.24
        },
        {
            town: "CLEMENTI",
            flatType: "4 ROOM",
            resalePrice: 63042.30,
            percentageChange: -2.13
        },
        {
            town: "BISHAN",
            flatType: "4 ROOM",
            resalePrice: 45022.20,
            percentageChange: 0.12
        },
    ]


</script>

<div class="p-5">
    <!-- PERSONAL STATS -->
    <div class="mb-10">
        <!-- Value of Properties Owned -->
         <h1 class="text-3xl pl-4">${portfolioValue.toFixed(2)}</h1>
         <!-- Graph -->
    </div>
    
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
            <div class="flex justify-between hover:bg-zinc-900 p-2 px-4 rounded-md">
                <div>
                    <p>{asset.town}</p>
                    <p class="text-sm italic">{asset.flatType}</p>
                </div>

                <div class="flex flex-col items-end">
                    <p>${asset.resalePrice.toFixed(2)}</p>
                    <p class="text-sm p-1 text-end w-[70px] rounded-lg {asset.percentageChange > 0 ? "bg-green-800 text-green-300" : "bg-red-800 text-red-300"}">{@html asset.percentageChange > 0 ? "&#8599;" : "&#x2198;"} {asset.percentageChange.toFixed(2)}%</p>
                </div>
            </div>
         {/each}

     </div>
    

     <div>
         <h1 class="text-xl">Leaderboards</h1>
         <!-- Top gainers, Top losers, Most listings -->
     </div>
</div>