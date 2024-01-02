<script lang="ts">
	export let value: string;
	export let id: string;
	export let text: string;
</script>

<div class="flex gap-3">
	<input on:click type="checkbox" {id} name="" {value} class="h-7 w-7 accent-PUA-dark-orange" />
	<label for={id}>
		<div class="truncate text-center text-xl font-bold leading-relaxed text-PUA-dark-red">
			{text}
		</div>
	</label>
</div>
