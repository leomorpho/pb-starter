import PocketBase from 'pocketbase';

// Initialize PocketBase client
export const pb = new PocketBase('http://localhost:8090');

// Enable auto-refresh for auth store
pb.autoCancellation(false);

export default pb;
