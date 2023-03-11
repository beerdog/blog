import { Component, createSignal, For, onMount } from 'solid-js'
import { IMetadata } from './interfaces';
import './BlogListView.scss'
import { useNavigate } from 'solid-app-router';

const BlogListView: Component = () => {
  const navigate = useNavigate()
    const [blogPosts, setBlogPosts] = createSignal<IMetadata[]>([]);
    onMount(async () => {
        const res = await fetch(`http://localhost:3001/api/metadata`);
        setBlogPosts(await res.json());
      });
      
  return (
    <ul id="blog-posts">
      <For each={blogPosts()} fallback={<p>Loading...</p>}>{ blogPost =>
        <div class="post">
          <a onclick={()=>{navigate(`/post/${blogPost.key}`)}}>{blogPost.title}</a>
          <div class="preamble">{blogPost.preamble}</div>
          <div class="footer">
            <div class="date-category">Posted at {blogPost.publishDate} in {blogPost.category}</div>
            <div class="tags">{blogPost.tags.map(t => `#${t} `)}</div>
          </div>
        </div>
      }</For>
    </ul>
  );
};

export default BlogListView;
