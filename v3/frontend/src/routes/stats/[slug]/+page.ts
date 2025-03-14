export function load({ params }) {
    let asset = params.slug
    let assetSplit = asset.split("+")
    let town = assetSplit[0]
    let flatType = assetSplit[1]

    console.log("TOWN: ", town)
    console.log("FLATTYPE: ", flatType)

    return {
        town,
        flatType,
    }
}