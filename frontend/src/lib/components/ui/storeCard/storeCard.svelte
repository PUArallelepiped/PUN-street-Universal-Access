<script lang="ts">
  import { page } from '$app/stores';
	import * as Card from '$lib/components/ui/card';

  export let name: string = "im pasta";
  export let description: string = "you are not pasta";
  export let picture: string = "https://i.imgur.com/T5zyE63.png";
  export let address: string = "pun street";
  export let rate: number = 0;
</script>

<a href="{$page.route.id}/{name}">
  <Card.Root class="relative w-[399px] h-[385px]">
    <Card.Content>
      <img src={picture} alt="" class="rounded w-[399px] h-[170px] mt-[10px]">
      <div class="w-[72px] h-[29px] bg-PUA-red rounded-full border-white border-2 border-opacity-50 absolute top-5 right-8">
        <div class="text-white text-base text-center font-semibold font-['Inter'] leading-tight my-1">{rate}</div>
      </div>
    </Card.Content>
    <Card.Header class="text-center">
      <Card.Title>{name}</Card.Title>
      <Card.Description>{description}</Card.Description>
    </Card.Header>
    <Card.Footer>
      <p>{address}</p>
    </Card.Footer>
  </Card.Root>
</a>