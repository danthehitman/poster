select 
ra.resource_id, user_id
from api.resource_authorizations as ra
union all
select rg.resource_id, ra.user_id
from api.resource_authorizations as ra
right join api.resource_groups as rg on rg.parent_resource_id = ra.resource_id