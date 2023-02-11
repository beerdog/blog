import { Component, createSignal, For, onMount } from 'solid-js'
import './BlogListView.scss'

const BlogListView: Component = () => {
    const [blogPosts, setBlogPosts] = createSignal([]);
    onMount(async () => {
        const res = await fetch(`http://localhost:3001/api/metadata`);
        setBlogPosts(await res.json());
      });
      
  return (
    
    <ul>
      <For each={blogPosts()} fallback={<p>Loading...</p>}>{ blogPost =>
        <div>{blogPost.title}</div>
      }</For>
    </ul>
  );
};

export default BlogListView;
