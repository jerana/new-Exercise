
/**
 * Definition for singly-linked list with a random pointer.
 * struct RandomListNode {
 *     int label;
 *     struct RandomListNode *next;
 *     struct RandomListNode *random;
 * };
 */

struct RandomListNode *copyRandomList(struct RandomListNode *head) {
    if (!head) {
        return head;
    }
    struct RandomListNode *copy , *start,*tmp;
    //Add node in between in list 
  start = head;
  while(start) {
      tmp = start->next;
      copy = (struct RandomListNode *)malloc(sizeof(struct RandomListNode));
      memset(copy,0,sizeof(struct RandomListNode));
      copy->label = start->label;
      start->next = copy;
      copy->next = tmp;
      start = tmp;
  }
  //Till this place,copy nodes puts in between original list nodes
  
  //Set Random pointers of Copy list 
  start = head;
  while(start) {
      if (start->random) {
          start->next->random = start->random->next;
      }
      start = start->next->next;
      
}
//Deattach copy list from original list
start = head;
copy = start->next;
struct RandomListNode * duplicate = copy;

while(start && copy) {
   
    tmp = copy->next;
    start->next = tmp;
    start = tmp;
    if (start) {
        copy->next = start->next;
        copy = start->next;
    } else {
        copy->next = start;
       
        copy = start;
    }
}
copy = duplicate;
tmp = head;
while(copy && tmp) {
    copy = copy->next;
    tmp = tmp->next;
}
return duplicate;
}

