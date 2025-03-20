export async function load({ params }) {
    const env = import.meta.env.VITE_ENV

    const url = env === "dev" ? "http://localhost:4000/frontpage?username=bed" : "https://hdbbuddy-production.up.railway.app/frontpage?username=bed"

    console.log("ENV: ", env, "\nURL: ", url)

    try {
        const response = await fetch(url)
        const data = await response.json()
        const { graphDataPoints, user, watchlistGraphDataPoints }: { graphDataPoints: GraphDataPoint[], user: any, watchlistGraphDataPoints:  GraphDataPoint[][]} = await data
        console.log(user, watchlistGraphDataPoints)
        return {
            user,
            graphDataPoints,
            watchlistGraphDataPoints,
        }

    } catch (e) {
        console.log("error detected, using mocked data...")
        console.error(e)
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
            ],
            watchlistGraphDataPoints: [
                [
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
                ],
                [
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
                ],
                [
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
                ],
            ]
        }
    }
}