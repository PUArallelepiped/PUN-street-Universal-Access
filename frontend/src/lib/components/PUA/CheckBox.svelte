<script lang="ts">
	export let value: string;
	export let id: string;
	export let text: string;
	export let onclick: () => void = () => {};
</script>

<div class="flex gap-3">
	<input
		on:click={onclick}
		type="checkbox"
		{id}
		name=""
		{value}
		class="h-7 w-7 accent-PUA-dark-orange"
	/>
	<label for={id}>
		<div class="text-center text-xl font-bold leading-relaxed text-PUA-dark-orange">
			{text}
		</div>
	</label>
</div>
