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