// See https://svelte.dev/docs/kit/types#app.d.ts
// for information about these interfaces
declare global {
	namespace App {
		// interface Error {}
		// interface Locals {}
		// interface PageData {}
		// interface PageState {}
		// interface Platform {}
	}
}

interface Asset {
	town: string
	flatType: string
	graph: DataSet[]
	resalePrice: number
	percentageChange: number
}


interface DataSet {
	label: string;
	data: (number | null)[];
	xAxis: string[];
	borderColor: string;
	backgroundColor: string;
}

interface GraphDataPoint {
	pricePerArea: number
	resalePrice: number
	date: string
}