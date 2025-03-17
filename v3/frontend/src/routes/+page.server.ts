export async function load({ params }) {
    const url = "http://localhost:4000/frontpage?username=bed"

    const response = await fetch(url)
    const data = await response.json()
    const { graphDataPoints, user }: { graphDataPoints: GraphDataPoint[], user: any} = await data

    return {
        user,
        graphDataPoints
    }
}