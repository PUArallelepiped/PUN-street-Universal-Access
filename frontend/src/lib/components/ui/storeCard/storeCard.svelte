<script lang="ts">
  import { page } from '$app/stores';
	import * as Card from '$lib/components/ui/card';

  export let name: string = "";
  export let description: string = "";
  export let picture: string = "";
  export let address: string = "";
</script>

<a href="{$page.route.id}/{name}">
  <Card.Root>
    <Card.Header class="text-center">
      <Card.Title>{name}</Card.Title>
      <Card.Description>{description}</Card.Description>
    </Card.Header>
    <Card.Content>
      <img src={picture} alt="" class="h-full object-cover object-center justify-center">
    </Card.Content>
    <Card.Footer>
      <p>{address}</p>
    </Card.Footer>
  </Card.Root>
</a>