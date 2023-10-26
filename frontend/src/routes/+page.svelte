<script lang="ts">
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

<h1>Welcome to PUA</h1>
<div class="bg-neutral-900">
	<h1 class="text-3xl font-bold underline">Hello world!</h1>
</div>

<p>here is root</p>

<div style="width: 40%; overflow: hidden;">
	<div class="text-2xl" style="float: left;">PUN Street Access</div>
</div>

<div style="display: flex; justify-content: flex-start;">
	<div style="width: 40%; float: left">
		<Input type="text" placeholder="QQㄋㄟㄋㄟ好喝到咩噗茶" class="max-w-xs" />

		<p style="padding-left: 10%; font-size: large;">Sort</p>
		<div>
			{#each sortCheckboxes as { id }}
				<div>
					<Checkbox {id} />
					<Label for={id}>{id}</Label>
				</div>
			{/each}
		</div>

		<p style="padding-left: 10%; font-size: large;">Price</p>
		<div style="display:flex; padding-right:20%">
			<div style="flex:1">NT$1</div>
			<div style="flex:1">NT$700</div>
			<div style="flex:1">NT$1000</div>
		</div>
		<div style="padding-right: 40%;">
			<DoubleRangeSlider bind:start bind:end />
		</div>

		<p style="padding-left: 10%; font-size: large;">Tag</p>
		<div>
			<div style="max-width: 60%;">
				{#each tagCheckboxes as { id }}
					<Checkbox {id} />
					<Label for={id}>{id}</Label>
				{/each}
			</div>
		</div>
	</div>

	<div style="display: flex-wrap; max-width: 70%">
		{#each cards as { title, description, content, footer }}
			<div style="display: flex-fill;min-width: 500px; float: left">
				<Card.Root>
					<Card.Header style="text-align:center">
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
<p>Visit <a href="./food" class="text-blue-400px"> Shop list</a></p>
