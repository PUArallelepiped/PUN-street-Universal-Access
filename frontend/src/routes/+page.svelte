<script lang="ts">
	//import * as Menubar from "$lib/components/ui/menubar";
	import * as Card from '$lib/components/ui/card';
	//import * as Select from "$lib/components/ui/select";
	import { Input } from '$lib/components/ui/input';
	import { Checkbox } from '$lib/components/ui/checkbox';
	import { Label } from '$lib/components/ui/label';
	import { Slider } from '$lib/components/ui/slider';
	//import { Button } from "$lib/components/ui/button";
	//let checked = false;
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
			<div>
				<Checkbox id="nice-shop" />
				<Label for="nice-shop">Nice shop</Label>
			</div>
			<div>
				<Checkbox id="fast-shop" />
				<Label for="fast-shop">Fast shop</Label>
			</div>
			<div>
				<Checkbox id="something-nice-for-you" />
				<Label for="something-nice-for-you">Something nice for you</Label>
			</div>
			<div>
				<Checkbox id="distance" />
				<Label for="distance">Distance</Label>
			</div>
		</div>

		<p style="padding-left: 10%; font-size: large;">Price</p>
		<div>
			<p>
				NT$1 &emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;NT$700
				&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;NT$1000
			</p>
			<Slider value={[50]} max={100} step={1} class="max-w-[85%]" />
		</div>

		<p style="padding-left: 10%; font-size: large;">Tag</p>
		<div>
			<div>
				<div>
					<Checkbox id="fresh" />
					<Label for="fresh">Fresh</Label>
					<Checkbox id="organic" />
					<Label for="organic">Organic</Label>
					<Checkbox id="healthy" />
					<Label for="healthy">Healthy</Label>
					<Checkbox id="dairy-free" />
					<Label for="dairy-free">Dairy-Free</Label>
				</div>
				<div>
					<Checkbox id="bakery" />
					<Label for="bakery">Bakery</Label>
					<Checkbox id="cafe" />
					<Label for="cafe">Cafe</Label>
					<Checkbox id="sugar-free" />
					<Label for="sugar-free">Sugar-Free</Label>
				</div>
				<div>
					<Checkbox id="vegetarian" />
					<Label for="vegetarian">Vegetarian</Label>
					<Checkbox id="free-delivery" />
					<Label for="free-delivery">Free-Delivery</Label>
				</div>
			</div>
		</div>
	</div>

	<div style="display: flex-wrap; float:left; max-width: 70%">
		<div style="float: left; flex: auto;min-width: 350px">
			<Card.Root>
				<Card.Header>
					<Card.Title>Card Title</Card.Title>
					<Card.Description>Card Description</Card.Description>
				</Card.Header>
				<Card.Content>
					<p>Card Content</p>
				</Card.Content>
				<Card.Footer>
					<p>Card Footer</p>
				</Card.Footer>
			</Card.Root>
		</div>
		<div style="float: left; flex: auto;min-width: 350px">
			<Card.Root>
				<Card.Header>
					<Card.Title>Card Title</Card.Title>
					<Card.Description>Card Description</Card.Description>
				</Card.Header>
				<Card.Content>
					<p>Card Content</p>
				</Card.Content>
				<Card.Footer>
					<p>Card Footer</p>
				</Card.Footer>
			</Card.Root>
		</div>
		<div style="float: left; flex: auto;min-width: 350px">
			<Card.Root>
				<Card.Header>
					<Card.Title>Card Title</Card.Title>
					<Card.Description>Card Description</Card.Description>
				</Card.Header>
				<Card.Content>
					<p>Card Content</p>
				</Card.Content>
				<Card.Footer>
					<p>Card Footer</p>
				</Card.Footer>
			</Card.Root>
		</div>
		<div style="float: left; flex: auto;min-width: 350px">
			<Card.Root>
				<Card.Header>
					<Card.Title>Card Title</Card.Title>
					<Card.Description>Card Description</Card.Description>
				</Card.Header>
				<Card.Content>
					<p>Card Content</p>
				</Card.Content>
				<Card.Footer>
					<p>Card Footer</p>
				</Card.Footer>
			</Card.Root>
		</div>
		<div style="float: left; flex: auto;min-width: 350px">
			<Card.Root>
				<Card.Header>
					<Card.Title>Card Title</Card.Title>
					<Card.Description>Card Description</Card.Description>
				</Card.Header>
				<Card.Content>
					<p>Card Content</p>
				</Card.Content>
				<Card.Footer>
					<p>Card Footer</p>
				</Card.Footer>
			</Card.Root>
		</div>
		<div style="float: left; flex: auto;min-width: 350px">
			<Card.Root>
				<Card.Header>
					<Card.Title>Card Title</Card.Title>
					<Card.Description>Card Description</Card.Description>
				</Card.Header>
				<Card.Content>
					<p>Card Content</p>
				</Card.Content>
				<Card.Footer>
					<p>Card Footer</p>
				</Card.Footer>
			</Card.Root>
		</div>
		<div style="float: left; flex: auto;min-width: 350px">
			<Card.Root>
				<Card.Header>
					<Card.Title>Card Title</Card.Title>
					<Card.Description>Card Description</Card.Description>
				</Card.Header>
				<Card.Content>
					<p>Card Content</p>
				</Card.Content>
				<Card.Footer>
					<p>Card Footer</p>
				</Card.Footer>
			</Card.Root>
		</div>
		<div style="float: left; flex: auto;min-width: 350px">
			<Card.Root>
				<Card.Header>
					<Card.Title>Card Title</Card.Title>
					<Card.Description>Card Description</Card.Description>
				</Card.Header>
				<Card.Content>
					<p>Card Content</p>
				</Card.Content>
				<Card.Footer>
					<p>Card Footer</p>
				</Card.Footer>
			</Card.Root>
		</div>
		<div style="float: left; flex: auto;min-width: 350px">
			<Card.Root>
				<Card.Header>
					<Card.Title>Card Title</Card.Title>
					<Card.Description>Card Description</Card.Description>
				</Card.Header>
				<Card.Content>
					<p>Card Content</p>
				</Card.Content>
				<Card.Footer>
					<p>Card Footer</p>
				</Card.Footer>
			</Card.Root>
		</div>
	</div>
</div>
<p>Visit <a href="./food" class="text-blue-400px"> Shop list</a></p>
