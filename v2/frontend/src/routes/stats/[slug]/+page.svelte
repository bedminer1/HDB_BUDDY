<script lang="ts">
	import LineChart from "$lib/components/LineChart.svelte";
	import * as Form from "$lib/components/ui/form";
	import { Input } from "$lib/components/ui/input";
	import { formSchema } from "./schema";
	import { superForm } from "sveltekit-superforms";
	import { toast } from "svelte-sonner";
	import { zodClient } from "sveltekit-superforms/adapters";

	export let data

	const form = superForm(data.form, {
		validators: zodClient(formSchema),
		onUpdated: ({ form: f }) => {
            console.log("Form updated:", f); // Debugging
            if (f.valid) {
                toast.success("Successfully submitted")
            } else {
                if ($message) {
                    toast.error($message)
                } else {
                    toast.error("Please fix the errors in the form.");
                }
            }
        }
	})

	const { form: formData, message, enhance } = form

	let dates: string[] = []
	$: {
		dates = []
		// @ts-ignore
		const recordWithMostRecords = data.records.reduce((maxRecord, current) =>
			current.records.length > maxRecord.records.length ? current : maxRecord
		)
		for (let record of recordWithMostRecords.records) {
			const recordDate = record.time.substring(0, 7);
			if (recordDate < $formData.start) {
				continue
			}
			if (recordDate > $formData.end) {
				break
			}
			dates.push(recordDate)
		}
	}

	let pricesArr: (number | null)[][] = []
	$: {
		// @ts-ignore
		pricesArr = data.records.map(townRecord => {
			let townPrices = Array(dates.length).fill(null);
			for (let timeRecord of townRecord.records) {
				const dateIndex = dates.indexOf(timeRecord.time.substring(0, 7));
				if (dateIndex !== -1) {
					townPrices[dateIndex] = timeRecord.averageResalePrice;
				}
			}
			return townPrices;
		})
	}

	function generateColors(count: number) {
		const colors = []
		for (let i = 0; i < count; i++) {
			const hue = (i * 360) / count // even distribute hues
			const color = {
				borderColor: `hsl(${hue}, 70%, 50%)`,
				backgroundColor: `hsl(${hue}, 70%, 60%)`,
			}
			colors.push(color)
		}
		return colors
	}

	$: colors = generateColors(data.records.length)
	let generatedObjects: DataSet[] = []
	$: {
		generatedObjects = data.records.map((record, i) => ({
			label: record.town,
			data: pricesArr[i],
			xAxis: dates,
			borderColor: colors[i].borderColor,
			backgroundColor: colors[i].backgroundColor
		}))
	}
</script>


<div class="flex flex-col justify-start items-center h-[95vh] w-full">
	<div class="w-full p-4">
		<LineChart
		{...{
			stats: generatedObjects,
			label: "Price(SGD)"
		}} />		
	</div>

	<form method="POST" use:enhance class="p-2 flex gap-4">
		<Form.Field {form} name="start">
			<Form.Control let:attrs>
			<Form.Label>Start Date</Form.Label>
			<Input {...attrs} bind:value={$formData.start} />
			</Form.Control>
			<Form.Description />
			<Form.FieldErrors />
		</Form.Field>
		<Form.Field {form} name="end">
			<Form.Control let:attrs>
				<Form.Label>End Date</Form.Label>
				<Input {...attrs} bind:value={$formData.end} />
			</Form.Control>
			<Form.Description />
			<Form.FieldErrors />
		</Form.Field>
	</form>
</div>