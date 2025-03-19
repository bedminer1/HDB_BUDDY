export async function load({ params }) {
    const url = "https://hdbbuddy-production.up.railway.app/frontpage?username=bed"

    try {
        const response = await fetch(url)
        const data = await response.json()
        const { graphDataPoints, user }: { graphDataPoints: GraphDataPoint[], user: any} = await data
    
        return {
            user,
            graphDataPoints
        }

    } catch (e) {
        console.log("error detected, using mocked data...")
        return {
            user: {
                mortgage: 202202,
            },
            graphDataPoints: [
                {
                    resalePrice: 500000,
                    pricePerArea: 2000,
                    date: "20 MAR 2020"
                },
                {
                    resalePrice: 520000,
                    pricePerArea: 2000,
                    date: "21 MAR 2020"
                },
                {
                    resalePrice: 560000,
                    pricePerArea: 2000,
                    date: "22 MAR 2020"
                },
                {
                    resalePrice: 470000,
                    pricePerArea: 2000,
                    date: "23 MAR 2020"
                },
            ]
        }
    }
}