<script lang="ts">
	export let name: string = '';
	export let id: string = '';
	export let checked: boolean = false;
	export let value: boolean | string = ' ';
	export let group: boolean | string | number = false;
</script>

<div class="flex">
	<div class="mb-[0.125rem] block min-h-[1.5rem] pl-[1.5rem]">
		<input
			class="dark:checked:border-primarydark:checked:after:border-PUA-dark-red relative float-left
			-ml-[1.5rem] mr-1 mt-0.5
			h-5 w-5
			appearance-none rounded-full border-2 border-solid border-PUA-dark-red before:pointer-events-none before:absolute before:h-4 before:w-4 before:scale-0
			before:rounded-full before:bg-transparent before:opacity-0 before:shadow-[0px_0px_0px_13px_transparent] before:content-[''] after:absolute after:z-[1] after:block
			after:h-4 after:w-4 after:rounded-full after:content-[''] checked:border-PUA-dark-red
			checked:before:opacity-[0.16] checked:after:absolute checked:after:left-1/2 checked:after:top-1/2 checked:after:h-[0.625rem] checked:after:w-[0.625rem] checked:after:rounded-full checked:after:border-primary
			checked:after:bg-PUA-dark-red checked:after:content-[''] checked:after:[transform:translate(-50%,-50%)] hover:cursor-pointer hover:before:opacity-[0.04]
			focus:shadow-none focus:outline-none
			focus:ring-0 focus:before:scale-100 focus:before:opacity-[0.12] focus:before:transition-[box-shadow_0.2s,transform_0.2s]
			checked:focus:border-PUA-dark-red checked:focus:before:transition-[box-shadow_0.2s,transform_0.2s] dark:border-PUA-dark-red
			dark:checked:after:bg-PUA-dark-red
			dark:checked:focus:border-primary
			dark:checked:focus:before:shadow-[0px_0px_0px_13px_#800000]"
			type="radio"
			{id}
			{name}
			{checked}
			{value}
			bind:group
		/>
		<label for={id}></label>
	</div>
</div>
