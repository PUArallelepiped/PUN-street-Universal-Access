<script lang="ts">
	import { page } from '$app/stores';

	//import * as Menubar from "$lib/components/ui/menubar";
	import * as Card from '$lib/components/ui/card';
	//import * as Select from "$lib/components/ui/select";
	import { Input } from '$lib/components/ui/input';
	import { Checkbox } from '$lib/components/ui/checkbox';
	import { Label } from '$lib/components/ui/label';
	//import { Slider } from '$lib/components/ui/slider';
	//import { Button } from "$lib/components/ui/button";
	//let checked = false;

	import DoubleRangeSlider from '$lib/components/ui/doubleSlider/doubleRangeSlider.svelte';
	let start = 0;
	let end = 1;

	let sortCheckboxes = [
		{ id: 'Nice shop' },
		{ id: 'Fast shop' },
		{ id: 'Something nice for you' },
		{ id: 'Distance' }
	];

	let tagCheckboxes = [
		{ id: 'Fresh' },
		{ id: 'Organic' },
		{ id: 'Healthy' },
		{ id: 'Dairy-free' },
		{ id: 'Bakery' },
		{ id: 'Cafe' },
		{ id: 'Sugar-free' },
		{ id: 'Vegetarian' },
		{ id: 'Free-delivery' }
	];

	let cards = [
		{ title: '1', description: 'aaa', content: 'aaa', footer: 'foota' },
		{ title: '2', description: 'bbb', content: 'bbb', footer: 'footb' },
		{ title: '3', description: 'ccc', content: 'ccc', footer: 'footc' },
		{ title: '4', description: 'ddd', content: 'ddd', footer: 'footd' },
		{ title: '5', description: 'eee', content: 'eee', footer: 'foote' },
		{ title: '6', description: 'fff', content: 'fff', footer: 'footf' }
	];
</script>

<div class="w-[40%] h-[50px] overflow-hidden">
	<div class="text-2xl float-left">PUN Street Access</div>
</div>

<div class="flex justify-start">
	<div class="w-[40%] float-left">
		<Input type="text" placeholder="QQㄋㄟㄋㄟ好喝到咩噗茶" class="max-w-xs" />

		<p class="text-center text-lg w-[325px] h-[30px]">Sort</p>
		<div>
			{#each sortCheckboxes as { id }}
				<div class="h-[30px]">
					<Checkbox {id} />
					<Label for={id}>{id}</Label>
				</div>
			{/each}
		</div>

		<p class="text-center text-lg w-[325px] h-[30px]">Price</p>
		<div class="flex justify-between w-[325px]">
			<div>NT$1</div>
			<div>NT$700</div>
			<div>NT$1000</div>
		</div>
		<div class="w-[325px] h-[30px]">
			<DoubleRangeSlider bind:start bind:end />
		</div>

		<p class="text-center text-lg w-[325px] h-[30px]">Tag</p>
		<div>
			<div class="max-w-60 pr-40">
				{#each tagCheckboxes as { id }}
					<div class="flex-wrap">
						<Checkbox {id} />
						<Label for={id}>{id}</Label>
					</div>
				{/each}
			</div>
		</div>
	</div>

	<div class="flex flex-wrap pl-20">
		{#each cards as { title, description, content, footer }}
			<div class="flex-fill w-[500px] float-left">
				<Card.Root>
					<Card.Header class="text-center">
						<Card.Title>{title}</Card.Title>
						<Card.Description>{description}</Card.Description>
					</Card.Header>
					<Card.Content>
						<p>{content}</p>
					</Card.Content>
					<Card.Footer>
						<p>{footer}</p>
					</Card.Footer>
				</Card.Root>
			</div>
		{/each}
	</div>
</div>

<h1>here is shop list</h1>
<p>Visit <a href="{$page.route.id}/impasta" class="text-blue-400"> Shop </a></p>
