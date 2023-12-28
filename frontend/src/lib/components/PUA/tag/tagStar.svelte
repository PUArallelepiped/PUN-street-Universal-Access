<script lang="ts">
	import star from '$lib/assets/star_2.svg';
	export let text: number | string = 0;
</script>

<div
	class=" flex h-7 w-16 items-center justify-center gap-1 rounded-[10px] border-4 border-red-200/40 bg-PUA-dark-red px-3 text-center text-base font-bold text-white"
>
	<img src={star} alt="" />
	<p class="mt-[1px]">{text}</p>
</div>
