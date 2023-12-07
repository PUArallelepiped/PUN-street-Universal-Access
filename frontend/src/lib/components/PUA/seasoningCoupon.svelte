<script lang="ts">
	async function changeseDate() {
		sTime = await changeDate(discount_start_date);
		eTime = await changeDate(discount_end_date);
	}
	async function changeDate(time: string) {
		let d = new Date(time);
		let month = '' + (d.getMonth() + 1),
			day = '' + d.getDate(),
			year = d.getFullYear();

		if (month.length < 2) month = '0' + month;
		if (day.length < 2) day = '0' + day;

		return [year, month, day].join('-');
	}
	export let used: boolean = true;
	export let name: string;
	export let percentage: number;
	export let discount_end_date: string;
	export let discount_start_date: string;
	let eTime: string;
	let sTime: string;
</script>

{#await changeseDate()}
	loading
{:then color}
	<div class="flex h-32 w-96 flex-col">
		<div
			class:text-PUA-gray={used}
			class:text-PUA-stone={!used}
			class:bg-PUA-stone={used}
			class:bg-PUA-gray={!used}
			class="
			 flex h-full items-center rounded-xl border-2 border-PUA-stone text-center font-semibold"
		>
			<div class=" px-6 py-5 text-base">{name}</div>
			<div class:bg-PUA-stone={!used} class:bg-PUA-gray={used} class=" h-12 w-1"></div>
			<div class="w-full">
				<div class="text-3xl">{percentage}%</div>
				<div>{sTime}~{eTime}</div>
			</div>
		</div>
	</div>
{/await}
