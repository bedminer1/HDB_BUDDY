<script lang="ts">
    import * as Select from "$lib/components/ui/select";

    // USER STATS
    let assets: Asset[] = []
    let portfolioValue: number = $state(0)


    let watchListedAssets: Asset[] = []


</script>

<div class="p-4">
    <!-- PERSONAL STATS -->
    <div class="mb-10">
        <!-- Value of Properties Owned -->
         <h1 class="text-3xl">${portfolioValue.toFixed(2)}</h1>
         <!-- Graph -->
    </div>
    
    <!-- WATCHLIST -->
     <div class="mb-10">
         <div class="flex justify-between">
            <h1>Prices</h1>
    
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

     </div>
    

    <h1>Leaderboards</h1>
    <!-- Top gainers, Top losers, Most listings -->
</div>