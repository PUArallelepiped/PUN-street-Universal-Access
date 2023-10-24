<h1>Welcome to PUA</h1>
<div class="bg-neutral-900">
	<h1 class="text-3xl font-bold underline">Hello world!</h1>
</div>

<p>here is root</p>

<script lang="ts">
  import * as Menubar from "$lib/components/ui/menubar";
  import { Input } from "$lib/components/ui/input";
	import { Checkbox } from "$lib/components/ui/checkbox";
  import { Label } from "$lib/components/ui/label";
	import { Slider } from "$lib/components/ui/slider";
  let checked = false;
</script>

<div style="width: 100%; overflow: hidden;">
  <div style="width: 1000px; float: left;">
		<p style="font-size: x-large; padding-left: 5%;">PUN Street Access</p>
  </div>
	<div style="float: right;">
		<Menubar.Root>
			<Menubar.Menu>
				<Menubar.Trigger>My Profile</Menubar.Trigger>
				<Menubar.Trigger>Shop List</Menubar.Trigger>
				<Menubar.Trigger>Check Order</Menubar.Trigger>
				<Menubar.Trigger>History</Menubar.Trigger>
				<Menubar.Trigger>Carts</Menubar.Trigger>
			</Menubar.Menu>
		</Menubar.Root>
  </div>
</div>
<div style="width: 100%; overflow: hidden;">
  <div style="width: 500px; float: left;"> 
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
			<p>NT$1 &emsp;&emsp;&emsp;&emsp;NT$700 &emsp;&emsp;&emsp;&emsp;NT$1000</p>
			<Slider value={[50]} max={100} step={1} class="max-w-[50%]" />
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
</div>
<p>Visit <a href="./food" class="text-blue-400"> Shop list</a></p>
